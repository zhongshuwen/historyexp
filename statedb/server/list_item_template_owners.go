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

func (srv *EOSServer) listItemTemplateOwnersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zlogger := logging.Logger(ctx, zlog)

	errors := validateListItemTemplateOwnersRequest(r)
	if len(errors) > 0 {
		writeError(ctx, w, derr.RequestValidationError(ctx, errors))
		return
	}

	request := extractListItemTemplateOwnersRequest(r)
	zlogger.Debug("extracted request", zap.Reflect("request", request))

	blockNum := uint64(request.BlockNum)
	actualBlockNum, _, _, speculativeWrites, err := srv.prepareRead(ctx, blockNum, false)
	if err != nil {
		writeError(ctx, w, fmt.Errorf("unable to prepare read: %w", err))
		return
	}
	zlogger.Debug("block_num and item id", zap.Uint64("block_num", blockNum),zap.Uint64("item_id", request.ItemId))

	tablet := statedb.NewItemTemplateOwnerTablet(request.ItemId)
	zlogger.Debug("got tablet in item template owners", zap.Reflect("tablet", tablet))

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

	zlogger.Debug("got tablet rows", zap.Reflect("tabletRows", tabletRows))
	zlogger.Debug("post-processing item template owners", zap.Int("item_owners_account_count", len(tabletRows)))
	itemTemplateOwners, err := sortedUniqueItemTemplateOwners(tabletRows)
	if err != nil {
		writeError(ctx, w, fmt.Errorf("unable to read unique item template owners tablet at %d: %w", blockNum, err))
		return
	}
	if len(itemTemplateOwners) == 0 {
		zlogger.Debug("no item template owners found for request, checking if we ever seen this public key")
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

	writeResponse(ctx, w, &listItemTemplateOwnersResponse{
		BlockNum:     actualBlockNum,
		ItemTemplateOwners:   itemTemplateOwners,
	})
}

type listItemTemplateOwnersRequest struct {
	ItemTemplateId    uint64 `json:"item_template_id"`
	BlockNum  uint64 `json:"block_num"`
}

type listItemTemplateOwnersResponse struct {
	BlockNum     uint64            `json:"block_num"`
	ItemTemplateOwners []*itemTemplateOwnerListItem `json:"item_owners"`
}
type itemTemplateOwnerListItem struct {
	ItemId    uint64 `json:"item_id"`
	Balance     uint64 `json:"balance"`
	AccountName string `json:"account_name"`
}

func validateListItemTemplateOwnersRequest(r *http.Request) url.Values {
	return validator.ValidateQueryParams(r, validator.Rules{
		"item_template_id":  []string{"required", "numeric_between:0,"},
		"block_num":  []string{"fluxdb.eos.blockNum"},
	})
}

func extractListItemTemplateOwnersRequest(r *http.Request) *listItemTemplateOwnersRequest {
	blockNum64, _ := strconv.ParseInt(r.FormValue("block_num"), 10, 64)
	itemTemplateId, _ := strconv.ParseInt(r.FormValue("item_template_id"), 10, 64)

	return &listItemTemplateOwnersRequest{
		ItemTemplateId: uint64(itemTemplateId),
		BlockNum:  uint64(blockNum64),
	}
}

var emptyItemTemplateOwners = []*itemTemplateOwnerListItem{}

func sortedUniqueItemTemplateOwners(tabletRows []fluxdb.TabletRow) ([]*itemTemplateOwnerListItem, error) {
	if len(tabletRows) <= 0 {
		// We return an actual array so the output is actually `[]` instead of `null`
		return emptyItemTemplateOwners, nil
	}




	out := make([]*itemTemplateOwnerListItem, len(tabletRows))
	i := 0
	for _, tabletRow := range tabletRows {
		itemId, owner := tabletRow.(*statedb.ItemTemplateOwnerRow).ItemIdAndOwner()

		balance, err := tabletRow.(*statedb.ItemTemplateOwnerRow).Balance()
		if err != nil {
			return emptyItemTemplateOwners, err
		}
		out[i] = &itemTemplateOwnerListItem{
			ItemId: itemId,
			Balance: accountNameSet[account],
			AccountName: account,
		}
		i++
	}


	sort.Slice(out, func(i, j int) bool {
		return out[i].ItemId < out[j].ItemId || (out[i].ItemId == out[j].ItemId && out[i].AccountName < out[j].AccountName)
	})

	return out, nil
}