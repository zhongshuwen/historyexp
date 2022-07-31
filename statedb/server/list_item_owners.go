// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	
	"github.com/dfuse-io/derr"
	"github.com/zhongshuwen/historyexp/statedb"
	"github.com/dfuse-io/fluxdb"
	"github.com/dfuse-io/logging"
	"github.com/dfuse-io/validator"
	"go.uber.org/zap"
)

func (srv *EOSServer) listItemOwnersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zlogger := logging.Logger(ctx, zlog)

	errors := validateListItemOwnersRequest(r)
	if len(errors) > 0 {
		writeError(ctx, w, derr.RequestValidationError(ctx, errors))
		return
	}

	request := extractListItemOwnersRequest(r)
	zlogger.Debug("extracted request", zap.Reflect("request", request))

	blockNum := uint64(request.BlockNum)
	actualBlockNum, _, _, speculativeWrites, err := srv.prepareRead(ctx, blockNum, false)
	if err != nil {
		writeError(ctx, w, fmt.Errorf("unable to prepare read: %w", err))
		return
	}

	tablet := statedb.NewItemOwnerTablet(request.ItemId)
	tabletRows, err := srv.db.ReadTabletAt(
		ctx,
		actualBlockNum,
		tablet,
		speculativeWrites,
	)
	if err != nil {
		writeError(ctx, w, fmt.Errorf("unable to read tablet at %d: %w", blockNum, err))
		return
	}

	zlogger.Debug("post-processing item owners", zap.Int("item_owners_account_count", len(tabletRows)))
	itemOwners, err := sortedUniqueItemOwners(tabletRows)
	if err != nil {
		writeError(ctx, w, fmt.Errorf("unable to read unique item owners tablet at %d: %w", blockNum, err))
		return
	}
	if len(itemOwners) == 0 {
		zlogger.Debug("no item owners found for request, checking if we ever seen this public key")
		seen, err := srv.db.HasSeenAnyRowForTablet(ctx, tablet)
		if err != nil {
			writeError(ctx, w, fmt.Errorf("unable to know if item id was seen once in db: %w", err))
			return
		}

		if !seen {
			writeError(ctx, w, statedb.DataItemIdNotFoundError(ctx, request.ItemId))
			return
		}
	}

	writeResponse(ctx, w, &listItemOwnersResponse{
		BlockNum:     actualBlockNum,
		ItemOwners:   itemOwners,
	})
}

type listItemOwnersRequest struct {
	ItemId    uint64 `json:"item_id"`
	BlockNum  uint64 `json:"block_num"`
}

type listItemOwnersResponse struct {
	BlockNum     uint64            `json:"block_num"`
	ItemOwners []*itemOwnerListItem `json:"item_owners"`
}
type itemOwnerListItem struct {
	Balance     uint64 `json:"block_num"`
	AccountName string `json:"account_name"`
}

func validateListItemOwnersRequest(r *http.Request) url.Values {
	return validator.ValidateQueryParams(r, validator.Rules{
		"item_id":  []string{"required", "numeric_between:0,"},
		"block_num":  []string{"fluxdb.eos.blockNum"},
	})
}

func extractListItemOwnersRequest(r *http.Request) *listItemOwnersRequest {
	blockNum64, _ := strconv.ParseInt(r.FormValue("block_num"), 10, 64)
	itemId, _ := strconv.ParseInt(r.FormValue("item_id"), 10, 64)

	return &listItemOwnersRequest{
		ItemId: uint64(itemId),
		BlockNum:  uint64(blockNum64),
	}
}

var emptyItemOwners = []*itemOwnerListItem{}

func sortedUniqueItemOwners(tabletRows []fluxdb.TabletRow) ([]*itemOwnerListItem, error) {
	if len(tabletRows) <= 0 {
		// We return an actual array so the output is actually `[]` instead of `null`
		return emptyItemOwners, nil
	}




	accountNameSet := map[string]uint64{}
	for _, tabletRow := range tabletRows {
		owner := tabletRow.(*statedb.ItemOwnerRow).Owner()

		balance, err := tabletRow.(*statedb.ItemOwnerRow).Balance()
		if err != nil {
			return emptyItemOwners, err
		}
		accountNameSet[owner] = balance
	}

	i := 0
	out := make([]*itemOwnerListItem, len(accountNameSet))
	for account := range accountNameSet {
		out[i] = &itemOwnerListItem{
			Balance: accountNameSet[account],
			AccountName: account,
		}
		i++
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].AccountName < out[j].AccountName
	})

	return out, nil
}
