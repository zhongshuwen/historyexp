package tokenmeta

import (
	"context"
	"encoding/json"
	"fmt"

	pbabicodec "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/abicodec/v1"
	pbtokenmeta "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/tokenmeta/v1"
	"github.com/zhongshuwen/historyexp/tokenmeta/cache"
zsw "github.com/zhongshuwen/zswchain-go"
	"go.uber.org/zap"
)

type accountsDbRow struct {
	Balance zsw.Asset `json:"balance"`
}

func (a *accountsDbRow) valid() bool {
	if a.Balance.Symbol.Symbol != "" {
		return true
	}
	return false
}

type EOSStakeDbRow struct {
	CPUWeight zsw.Asset       `json:"cpu_weight"`
	NetWeight zsw.Asset       `json:"net_weight"`
	To        zsw.AccountName `json:"to"`
	From      zsw.AccountName `json:"from"`
}

func (a *EOSStakeDbRow) valid() bool {
	if a.CPUWeight.Symbol.Symbol == "" {
		return false
	}
	if a.NetWeight.Symbol.Symbol == "" {
		return false
	}
	return true
}

type statDbRow struct {
	Issuer    zsw.AccountName `json:"issuer"`
	MaxSupply zsw.Asset       `json:"max_supply"`
	Supply    zsw.Asset       `json:"supply"`
}

func (s *statDbRow) valid() bool {
	if s.Issuer == "" {
		return false
	}
	if s.Supply.Symbol.Symbol == "" {
		return false
	}
	return true
}

type abiItem struct {
	abi      *zsw.ABI
	blockNum uint32
}

func (t *TokenMeta) getABI(contract zsw.AccountName, blockNum uint32) (*zsw.ABI, error) {
	if abiItem, ok := t.abisCache[string(contract)]; ok {
		return abiItem.abi, nil
	}

	zlog.Info("abi cache miss", zap.String("contract", string(contract)), zap.Uint32("at_block_num", blockNum))
	resp, err := t.abiCodecCli.GetAbi(context.Background(), &pbabicodec.GetAbiRequest{
		Account:    string(contract),
		AtBlockNum: blockNum,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get abi for contract %q: %w", string(contract), err)
	}

	var abi *zsw.ABI
	err = json.Unmarshal([]byte(resp.JsonPayload), &abi)
	if err != nil {
		return nil, fmt.Errorf("unable to decode abi for contract %q: %w", string(contract), err)
	}

	// store abi in cache for late uses
	t.abisCache[string(contract)] = &abiItem{
		abi:      abi,
		blockNum: resp.AbiBlockNum,
	}

	return abi, nil
}

func getStakeEntryFromDBRow(contract zsw.AccountName, scope string, dbRow json.RawMessage) (*cache.EOSStakeEntry, error) {
	stakeRow := &EOSStakeDbRow{}
	err := json.Unmarshal(dbRow, &stakeRow)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarsal EOS stake db row: %s", string(dbRow))
	}
	if !stakeRow.valid() {
		return nil, fmt.Errorf("invalid stake row: %s", string(dbRow))
	}
	if stakeRow.From != zsw.AccountName(scope) {
		zlog.Warn("failed assumption: EOS stake FROM is not == scope",
			zap.String("contract", string(contract)),
			zap.String("scope", scope),
			zap.String("from", string(stakeRow.From)),
		)
	}

	return &cache.EOSStakeEntry{
		From: stakeRow.From,
		To:   stakeRow.To,
		Net:  stakeRow.NetWeight.Amount,
		Cpu:  stakeRow.CPUWeight.Amount,
	}, nil
}

func getAccountBalanceFromDBRow(contract zsw.AccountName, symbol *zsw.Symbol, scope string, dbRow json.RawMessage) (*pbtokenmeta.AccountBalance, error) {
	accountRow := &accountsDbRow{}
	err := json.Unmarshal(dbRow, &accountRow)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarsal accounts table row: %s", string(dbRow))
	}

	if !accountRow.valid() {
		return nil, fmt.Errorf("invalid accounts row: %s", string(dbRow))
	}

	return &pbtokenmeta.AccountBalance{
		TokenContract: string(contract),
		Account:       scope,
		Amount:        uint64(accountRow.Balance.Amount),
		Precision:     uint32(accountRow.Balance.Precision),
		Symbol:        symbol.Symbol,
	}, nil
}

func getTokenFromDBRow(contract zsw.AccountName, symbol *zsw.Symbol, dbRow json.RawMessage) (*pbtokenmeta.Token, error) {
	statRow := &statDbRow{}
	err := json.Unmarshal(dbRow, &statRow)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarsal table row: %s", string(dbRow))
	}

	if symbol == nil {
		symbol = &zsw.Symbol{
			Precision: statRow.Supply.Symbol.Precision,
			Symbol:    statRow.Supply.Symbol.Symbol,
		}
	}

	if !statRow.valid() {
		return nil, fmt.Errorf("invalid stat row: %s", string(dbRow))
	}

	return &pbtokenmeta.Token{
		Contract:      string(contract),
		Symbol:        symbol.Symbol,
		Precision:     uint32(symbol.Precision),
		Issuer:        string(statRow.Issuer),
		MaximumSupply: uint64(statRow.MaxSupply.Amount),
		TotalSupply:   uint64(statRow.Supply.Amount),
		Holders:       0,
	}, nil
}

func decodeTableRow(data []byte, tableName zsw.TableName, abi *zsw.ABI) (json.RawMessage, error) {
	out, err := abi.DecodeTableRow(tableName, data)
	if err != nil {
		return nil, err
	}
	return out, err
}
