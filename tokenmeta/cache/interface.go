package cache

import (
	"time"

	"github.com/dfuse-io/bstream"
	pbtokenmeta "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/tokenmeta/v1"
zsw "github.com/zhongshuwen/zswchain-go"
)

type Cache interface {
	Tokens() []*pbtokenmeta.Token
	IsTokenContract(contract zsw.AccountName) bool
	TokenContract(contract zsw.AccountName, code zsw.SymbolCode) *pbtokenmeta.Token
	AccountBalances(account zsw.AccountName, opts ...AccountBalanceOption) []*OwnedAsset
	TokenBalances(contract zsw.AccountName, opts ...TokenBalanceOption) []*OwnedAsset
	Apply(mutationsBatch *MutationsBatch, processedBlock bstream.BlockRef) []error
	SaveToFile() error
	AtBlockRef() bstream.BlockRef
	SetHeadBlockTime(t time.Time)
	GetHeadBlockTime() time.Time
}

const EOSTokenContract = zsw.AccountName("zswhq.token")

type SortingOrder int32

const (
	ASC SortingOrder = iota
	DESC
)

type TokenBalanceOption int

const (
	EOSIncludeStakedTokOpt TokenBalanceOption = iota
)

type AccountBalanceOption int

const (
	EOSIncludeStakedAccOpt AccountBalanceOption = iota
)
