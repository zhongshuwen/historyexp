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

package statedb

import (
	"bytes"
	"fmt"
	"encoding/binary"
	"github.com/dfuse-io/bstream"
	pbcodec "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/codec/v1"
	"github.com/dfuse-io/fluxdb"
	"go.uber.org/zap"

)

type BlockMapper struct {
}


const (
	ZswItemsMintAction = 0xfffffff0
	ZswItemsTransferAction = 0xfffffff1
)

type TableItemBalancesRow []byte
func (tibr TableItemBalancesRow) ItemId() uint64 {
	return binary.LittleEndian.Uint64(tibr[0:8])
}
func (tibr TableItemBalancesRow) TotalBalance() uint64 {
	return binary.LittleEndian.Uint64(tibr[12:20]) + binary.LittleEndian.Uint64(tibr[20:28]) + binary.LittleEndian.Uint64(tibr[28:36])
}
func (m *BlockMapper) Map(rawBlk *bstream.Block) (*fluxdb.WriteRequest, error) {
	blk := rawBlk.ToNative().(*pbcodec.Block)

	lastSingletEntryMap := map[string]fluxdb.SingletEntry{}
	lastTabletRowMap := map[string]fluxdb.TabletRow{}

	firstDbOpWasInsert := map[string]bool{}
	
	itemsActTypeMap := map[uint32]uint32{}
	itemsOrdinalToExecutionIndexMap := map[uint32]uint32{}


	req := &fluxdb.WriteRequest{
		Height:   rawBlk.Num(),
		BlockRef: rawBlk.AsRef(),
	}

	blockNum := req.BlockRef.Num()
	for _, trx := range blk.TransactionTraces() {
		actionMatcher := blk.FilteringActionMatcher(trx, isRequiredSystemAction)
		for _, a := range trx.ActionTraces {
			if a.Account() == "zsw.items" {
				if a.Name() == "mint" && a.Receiver == "zsw.items" {
					itemsActTypeMap[a.ExecutionIndex] = ZswItemsMintAction
					itemsOrdinalToExecutionIndexMap[a.ActionOrdinal] = a.ExecutionIndex+1
				}else if a.Name() == "transfer" && a.Receiver == "zsw.items" {
					itemsActTypeMap[a.ExecutionIndex] = ZswItemsTransferAction
					itemsOrdinalToExecutionIndexMap[a.ActionOrdinal] = a.ExecutionIndex+1
				}else if a.Name() == "logtransfer" && a.Receiver == "zsw.items"{
					if itemsOrdinalToExecutionIndexMap[a.CreatorActionOrdinal] != 0 {
						itemsActTypeMap[itemsOrdinalToExecutionIndexMap[a.CreatorActionOrdinal+1]] = a.ExecutionIndex+1
					}
				}
			}
		}


		for _, dbOp := range trx.DbOps {
			if traceEnabled {
				zlog.Debug("db op", zap.Reflect("op", dbOp))
				if dbOp.Code == "zsw.items" && dbOp.TableName == "itembalances" {

					actItemsType2 := itemsActTypeMap[dbOp.ActionIndex]
					zlog.Debug("items op", 
						zap.String("dbOp.Code", dbOp.Code), 
						zap.String("dbOp.TableName", dbOp.TableName),
						zap.Uint32("dbOp.TableName", itemsActTypeMap[dbOp.ActionIndex]),
						
						zap.Bool("(actItemsType == ZswItemsMintAction || actItemsType == ZswItemsTransferAction)", (actItemsType2 == ZswItemsMintAction || actItemsType2 == ZswItemsTransferAction)),
						
						zap.Bool("itemsActTypeMap[dbOp.ActionIndex] != 0",itemsActTypeMap[dbOp.ActionIndex] != 0),
						
					)
					zlog.Debug("db op items", zap.Reflect("op", dbOp))
				}
			}
			actItemsType := itemsActTypeMap[dbOp.ActionIndex]
			if dbOp.Code == "zsw.items" && dbOp.TableName == "itembalances" && (actItemsType == ZswItemsMintAction || actItemsType == ZswItemsTransferAction) && itemsActTypeMap[dbOp.ActionIndex] != 0{

				zlog.Debug("db op items good", zap.Reflect("op", dbOp))
				//logActionIndex := itemsActTypeMap[dbOp.ActionIndex] - 1
				if dbOp.Operation == pbcodec.DBOp_OPERATION_UPDATE {
					if !bytes.Equal(dbOp.OldData, dbOp.NewData) {
						zlog.Debug("update_good", zap.Uint64("item_id", TableItemBalancesRow(dbOp.NewData).ItemId()), zap.Uint64("total_balance", TableItemBalancesRow(dbOp.NewData).TotalBalance()))
						itemOwnerRow, err := NewItemOwnerRow(blockNum, TableItemBalancesRow(dbOp.NewData).ItemId(), TableItemBalancesRow(dbOp.NewData).TotalBalance(), dbOp.Scope, false)
						if err != nil {
							return nil, fmt.Errorf("unable to extract item owner: %w", err)
						}

						zlog.Debug("db op items good update", zap.Reflect("op", itemOwnerRow))
						lastTabletRowMap[keyForRow(itemOwnerRow)] = itemOwnerRow
					}
				}else if dbOp.Operation == pbcodec.DBOp_OPERATION_REMOVE {
					itemOwnerRow, err := NewItemOwnerRow(blockNum, TableItemBalancesRow(dbOp.OldData).ItemId(), 0, dbOp.Scope, true)
					if err != nil {
						return nil, fmt.Errorf("unable to extract item owner: %w", err)
					}
					rowKey := keyForRow(itemOwnerRow)

					zlog.Debug("db op items good remove", zap.Reflect("op", itemOwnerRow))
					if firstDbOpWasInsert[rowKey] {
						delete(firstDbOpWasInsert, rowKey)
						delete(lastTabletRowMap, rowKey)
					} else {
						lastTabletRowMap[rowKey] = itemOwnerRow
					}
				}else if dbOp.Operation == pbcodec.DBOp_OPERATION_INSERT {

					zlog.Debug("insert_good", zap.Uint64("item_id", TableItemBalancesRow(dbOp.NewData).ItemId()), zap.Uint64("total_balance", TableItemBalancesRow(dbOp.NewData).TotalBalance()))
					itemOwnerRow, err := NewItemOwnerRow(blockNum, TableItemBalancesRow(dbOp.NewData).ItemId(),TableItemBalancesRow(dbOp.NewData).TotalBalance(), dbOp.Scope, false)
					if err != nil {
						return nil, fmt.Errorf("unable to extract item owner: %w", err)
					}
					rowKey := keyForRow(itemOwnerRow)
					lastOp := lastTabletRowMap[rowKey]
					if lastOp == nil {
						firstDbOpWasInsert[rowKey] = true
					}
					zlog.Debug("db op items good insert", zap.Reflect("op", itemOwnerRow))
					lastTabletRowMap[rowKey] = itemOwnerRow
				}
			}

			if !actionMatcher.Matched(dbOp.ActionIndex) {
				continue
			}

			// There is no change in this row, not sure how it got here, discarding it anyway
			if dbOp.Operation == pbcodec.DBOp_OPERATION_UPDATE && bytes.Equal(dbOp.OldData, dbOp.NewData) && dbOp.OldPayer == dbOp.NewPayer {
				continue
			}

			row, err := dbOpToContractStateRow(blockNum, dbOp)
			if err != nil {
				return nil, fmt.Errorf("unable to create contract state row for db op: %w", err)
			}

			rowKey := keyForRow(row)
			lastOp := lastTabletRowMap[rowKey]
			if lastOp == nil && dbOp.Operation == pbcodec.DBOp_OPERATION_INSERT {
				firstDbOpWasInsert[rowKey] = true
			}

			if dbOp.Operation == pbcodec.DBOp_OPERATION_REMOVE && firstDbOpWasInsert[rowKey] {
				delete(firstDbOpWasInsert, rowKey)
				delete(lastTabletRowMap, rowKey)
			} else {
				lastTabletRowMap[rowKey] = row
			}
		}

		// All perms ops comes from required system actions, so we process them all
		for _, permOp := range trx.PermOps {
			rows, err := permOpToKeyAccountRows(blockNum, permOp)
			if err != nil {
				return nil, fmt.Errorf("unable to create key account rows for perm op: %w", err)
			}

			for _, row := range rows {
				lastTabletRowMap[keyForRow(row)] = row
			}
		}

		for _, tableOp := range trx.TableOps {
			if !actionMatcher.Matched(tableOp.ActionIndex) {
				continue
			}

			row, err := NewContractTableScopeRow(blockNum, tableOp)
			if err != nil {
				return nil, fmt.Errorf("unable to create contract table scope row for table op: %w", err)
			}

			lastTabletRowMap[keyForRow(row)] = row
		}

		for _, act := range trx.ActionTraces {
			if act.Receiver != "zswhq" && act.Receiver != "zsw.items" {
				continue
			}

			// We always process those regardless of the filtering applied to the block since they are all system actions
			switch act.SimpleName() {
			case "zswhq:setabi":
				abiEntry, err := NewContractABIEntry(req.BlockRef.Num(), act)
				if err != nil {
					return nil, fmt.Errorf("unable to extract abi entry: %w", err)
				}

				if abiEntry == nil {
					zlog.Debug("abi entry not added since it was not decoded correctly")
					continue
				}

				lastSingletEntryMap[keyForEntry(abiEntry)] = abiEntry

			case "zswhq:linkauth":
				authLinkRow, err := NewInsertAuthLinkRow(blockNum, act)
				if err != nil {
					return nil, fmt.Errorf("unable to extract link auth: %w", err)
				}

				lastTabletRowMap[keyForRow(authLinkRow)] = authLinkRow

			case "zswhq:unlinkauth":
				authLinkRow, err := NewDeleteAuthLinkRow(blockNum, act)
				if err != nil {
					return nil, fmt.Errorf("unable to extract unlink auth: %w", err)
				}

				lastTabletRowMap[keyForRow(authLinkRow)] = authLinkRow
			}
		}
	}

	addSingletEntriesToRequest(req, lastSingletEntryMap)
	addTabletRowsToRequest(req, lastTabletRowMap)

	return req, nil
}

func isRequiredSystemAction(actTrace *pbcodec.ActionTrace) bool {
	if actTrace.Receiver != "zswhq" || actTrace.Action.Account != "zswhq" {
		return false
	}

	actionName := actTrace.Action.Name
	return actionName == "setabi" || actionName == "newaccount" || actionName == "updateauth" || actionName == "deleteauth" || actionName == "linkauth" || actionName == "unlinkauth"
}

func addSingletEntriesToRequest(request *fluxdb.WriteRequest, singleEntriesMap map[string]fluxdb.SingletEntry) {
	for _, entry := range singleEntriesMap {
		request.AppendSingletEntry(entry)
	}
}

func addTabletRowsToRequest(request *fluxdb.WriteRequest, tabletRowsMap map[string]fluxdb.TabletRow) {
	for _, row := range tabletRowsMap {
		request.AppendTabletRow(row)
	}
}

func addDBOpsToWriteRequest(request *fluxdb.WriteRequest, latestDbOps map[string]*pbcodec.DBOp) error {
	blockNum := request.BlockRef.Num()
	for _, op := range latestDbOps {
		row, err := NewContractStateRow(blockNum, op)
		if err != nil {
			return fmt.Errorf("unable to create row for db op: %w", err)
		}

		request.AppendTabletRow(row)
	}

	return nil
}

func dbOpToContractStateRow(blockNum uint64, op *pbcodec.DBOp) (*ContractStateRow, error) {
	row, err := NewContractStateRow(blockNum, op)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func permOpToKeyAccountRows(blockNum uint64, permOp *pbcodec.PermOp) ([]*KeyAccountRow, error) {
	switch permOp.Operation {
	case pbcodec.PermOp_OPERATION_INSERT:
		return permToKeyAccountRows(blockNum, permOp.NewPerm, false)
	case pbcodec.PermOp_OPERATION_UPDATE:
		var rows []*KeyAccountRow
		deletedRows, err := permToKeyAccountRows(blockNum, permOp.OldPerm, true)
		if err != nil {
			return nil, fmt.Errorf("unable to get key accounts from old perm: %w", err)
		}

		insertedRows, err := permToKeyAccountRows(blockNum, permOp.NewPerm, false)
		if err != nil {
			return nil, fmt.Errorf("unable to get key accounts from new perm: %w", err)
		}

		rows = append(rows, deletedRows...)
		rows = append(rows, insertedRows...)

		return rows, nil
	case pbcodec.PermOp_OPERATION_REMOVE:
		return permToKeyAccountRows(blockNum, permOp.OldPerm, true)
	}

	panic(fmt.Errorf("unknown perm op %s", permOp.Operation))
}

func permToKeyAccountRows(blockNum uint64, perm *pbcodec.PermissionObject, isDeletion bool) (rows []*KeyAccountRow, err error) {
	if perm.Authority == nil || len(perm.Authority.Keys) == 0 {
		return nil, nil
	}

	rows = make([]*KeyAccountRow, len(perm.Authority.Keys))
	for i, key := range perm.Authority.Keys {
		rows[i], err = NewKeyAccountRow(blockNum, key.PublicKey, perm.Owner, perm.Name, isDeletion)
		if err != nil {
			if err != nil {
				return nil, fmt.Errorf("unable to create key account row for permission object: %w", err)
			}
		}
	}

	return
}

func keyForEntry(entry fluxdb.SingletEntry) string {
	return string(fluxdb.KeyForSingletEntry(entry))
}

func keyForRow(row fluxdb.TabletRow) string {
	return string(fluxdb.KeyForTabletRow(row))
}
