package cache

import (
	"testing"

	"github.com/dfuse-io/bstream"
	pbtokenmeta "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/tokenmeta/v1"
zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultCache_AccountBalances(t *testing.T) {
	tests := []*struct {
		name              string
		accountName       zsw.AccountName
		balances          map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset
		eosStake          map[zsw.AccountName]*EOSStake
		expectOwnedAssets []*OwnedAsset
		options           []AccountBalanceOption
	}{
		{
			name:        "owner with one token in one contract",
			accountName: zsw.AccountName("eoscanadadad"),
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
					zsw.AccountName("johndoemyhero"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(29371, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(100, "EOS"),
						Contract: "zswhq.token",
					},
				},
			},
		},
		{
			name:        "owner with one token in two contract",
			accountName: zsw.AccountName("eoscanadadad"),
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
				zsw.AccountName("abababababa"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(200, "WALL"),
								Contract: "abababababa",
							},
						},
					},
				},
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(100, "EOS"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(200, "WALL"),
						Contract: "abababababa",
					},
				},
			},
		},
		{
			name:        "owner with two tokens in one contract",
			accountName: zsw.AccountName("eoscanadadad"),
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(3, "WALL"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(100, "EOS"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(3, "WALL"),
						Contract: "zswhq.token",
					},
				},
			},
		},
		{
			name:        "poor owner without any assets",
			accountName: zsw.AccountName("johndoemyone"),
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
		},
		{
			name:        "owner with staked EOS",
			accountName: zsw.AccountName("eoscanadadad"),
			eosStake: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("eoscanadadad"): {
					TotalNet: zsw.Int64(24),
					TotalCpu: zsw.Int64(14),
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("eoscanadadad"): {
							To:   "eoscanadadad",
							From: "eoscanadadad",
							Net:  zsw.Int64(24),
							Cpu:  zsw.Int64(14),
						},
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(111, "WAX"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			options: []AccountBalanceOption{
				EOSIncludeStakedAccOpt,
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(111, "WAX"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(138, "EOS"),
						Contract: "zswhq.token",
					},
				},
			},
		},
		{
			name:        "owner without staked EOS, requested",
			accountName: zsw.AccountName("eoscanadadad"),
			eosStake:    map[zsw.AccountName]*EOSStake{},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			options: []AccountBalanceOption{
				EOSIncludeStakedAccOpt,
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(100, "EOS"),
						Contract: "zswhq.token",
					},
				},
			},
		},
		{
			name:        "owner with unwanted staked EOS",
			accountName: zsw.AccountName("eoscanadadad"),
			eosStake: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("eoscanadadad"): {

					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("eoscanadadad"): {
							To:   "eoscanadad",
							From: "eoscanadadad",
							Net:  zsw.Int64(24),
							Cpu:  zsw.Int64(14),
						},
					},
				},
			},

			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(100, "EOS"),
						Contract: "zswhq.token",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				Balances: test.balances,
				EOSStake: test.eosStake,
			}
			ownedAssets := cache.AccountBalances(test.accountName, test.options...)
			assert.ElementsMatch(t, test.expectOwnedAssets, ownedAssets)
		})
	}
}

func TestDefaultCache_TokenBalances(t *testing.T) {
	tests := []*struct {
		name              string
		contract          zsw.AccountName
		balances          map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset
		expectOwnedAssets []*OwnedAsset
		eosStake          map[zsw.AccountName]*EOSStake
		options           []TokenBalanceOption
	}{
		{
			name:     "contract with multiple users and tokens",
			contract: zsw.AccountName("zswhq.token"),
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(3, "WALL"),
								Contract: "zswhq.token",
							},
						},
					},
					zsw.AccountName("johndoeonecoin"): {
						{
							Owner: zsw.AccountName("johndoeonecoin"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(23927, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(100, "EOS"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(3, "WALL"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("johndoeonecoin"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(23927, "EOS"),
						Contract: "zswhq.token",
					},
				},
			},
		},
		{
			name:     "zswhq.token with staked",
			contract: zsw.AccountName("zswhq.token"),
			options: []TokenBalanceOption{
				EOSIncludeStakedTokOpt,
			},
			eosStake: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("eoscanadadad"): {
					TotalCpu: zsw.Int64(200000000),
					TotalNet: zsw.Int64(100000000),
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("eoscanadadad"): {
							To:   "eoscanadad",
							From: "eoscanadadad",
							Net:  zsw.Int64(100000000),
							Cpu:  zsw.Int64(200000000),
						},
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(3, "WALL"),
								Contract: "zswhq.token",
							},
						},
					},
					zsw.AccountName("johndoeonecoin"): {
						{
							Owner: zsw.AccountName("johndoeonecoin"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(23927, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
			expectOwnedAssets: []*OwnedAsset{
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(300000100, "EOS"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("eoscanadadad"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(3, "WALL"),
						Contract: "zswhq.token",
					},
				},
				{
					Owner: zsw.AccountName("johndoeonecoin"),
					Asset: &zsw.ExtendedAsset{
						Asset:    generateTestAsset(23927, "EOS"),
						Contract: "zswhq.token",
					},
				},
			},
		},
		{
			name:     "contract does not exists",
			contract: zsw.AccountName("eidoeonecoin"),
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				zsw.AccountName("zswhq.token"): {
					zsw.AccountName("eoscanadadad"): {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: "zswhq.token",
							},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				Balances: test.balances,
				EOSStake: test.eosStake,
			}
			assert.ElementsMatch(t, test.expectOwnedAssets, cache.TokenBalances(test.contract, test.options...))
		})
	}
}

func TestDefaultCache_IsTokenContract(t *testing.T) {
	tests := []struct {
		name        string
		contract    zsw.AccountName
		tokens      map[zsw.AccountName][]*pbtokenmeta.Token
		expectValue bool
	}{
		{
			name:        "contract is not cached",
			contract:    zsw.AccountName("zswhq.token"),
			tokens:      map[zsw.AccountName][]*pbtokenmeta.Token{},
			expectValue: false,
		},
		{
			name:     "contract is cached",
			contract: zsw.AccountName("zswhq.token"),
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{},
			},
			expectValue: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				TokensInContract: test.tokens,
			}
			assert.Equal(t, test.expectValue, cache.IsTokenContract(test.contract))
		})
	}

}

func TestDefaultCache_hasSymbolForContract(t *testing.T) {
	tests := []struct {
		name        string
		contract    zsw.AccountName
		symbol      string
		tokens      map[zsw.AccountName][]*pbtokenmeta.Token
		expectValue bool
	}{
		{
			name:     "contract and symbol exists",
			contract: zsw.AccountName("zswhq.token"),
			symbol:   "EOS",
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol: "EOS",
					},
				},
			},
			expectValue: true,
		},
		{
			name:        "contract does not exists",
			contract:    zsw.AccountName("zswhq.token"),
			tokens:      map[zsw.AccountName][]*pbtokenmeta.Token{},
			expectValue: false,
		},
		{
			name:     "contract exists but symbol does not exists",
			contract: zsw.AccountName("zswhq.token"),
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol: "WAX",
					},
				},
			},
			expectValue: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				TokensInContract: test.tokens,
			}
			assert.Equal(t, test.expectValue, cache.hasSymbolForContract(test.contract, test.symbol))
		})
	}

}

func TestDefaultCache_setBalance(t *testing.T) {
	tests := []struct {
		name           string
		asset          *OwnedAsset
		tokens         map[zsw.AccountName][]*pbtokenmeta.Token
		balances       map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset
		expectBalances map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset
		expectTokens   map[zsw.AccountName][]*pbtokenmeta.Token
		expectError    bool
	}{
		{
			name: "sunny path",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(100, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol: "EOS",
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
		},
		{
			name: "sunny path when account already seen",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(100, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 0,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
		},
		{
			name: "set a new balance for a non existing contract",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(100, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens:         map[zsw.AccountName][]*pbtokenmeta.Token{},
			balances:       map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{},
			expectTokens:   map[zsw.AccountName][]*pbtokenmeta.Token{},
			expectError:    true,
		},
		{
			name: "set a new balance for a non-existing token",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(100, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{},
			},
			balances:       map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{},
			},
			expectError: true,
		},
		{
			name: "change an existing balance a new balance",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(100, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(20, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
		},
		{
			name: "should only change the specfic contract",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(100, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{Symbol: "EOS", Holders: 1},
				},
				zsw.AccountName("eidosonecoin"): []*pbtokenmeta.Token{
					{Symbol: "EOS", Holders: 1},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(20, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
				"eidosonecoin": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(30, "EOS"),
								Contract: zsw.AccountName("eidosonecoin"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
				"eidosonecoin": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(30, "EOS"),
								Contract: zsw.AccountName("eidosonecoin"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{Symbol: "EOS", Holders: 1},
				},
				zsw.AccountName("eidosonecoin"): []*pbtokenmeta.Token{
					{Symbol: "EOS", Holders: 1},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				TokensInContract: test.tokens,
				Balances:         test.balances,
			}
			err := cache.setBalance(test.asset)
			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, test.expectBalances, cache.Balances)
			assert.Equal(t, test.expectTokens, cache.TokensInContract)

		})
	}

}

func TestDefaultCache_removeBalance(t *testing.T) {
	tests := []struct {
		name           string
		asset          *OwnedAsset
		tokens         map[zsw.AccountName][]*pbtokenmeta.Token
		balances       map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset
		expectBalances map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset
		expectTokens   map[zsw.AccountName][]*pbtokenmeta.Token
		expectError    bool
	}{
		{
			name: "remove an existing balance",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(0, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 0,
					},
				},
			},
		},
		{
			name: "remove an existing balance while maintaining another token",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(0, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
					{
						Symbol:  "WALL",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(65, "WALL"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(65, "WALL"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 0,
					},
					{
						Symbol:  "WALL",
						Holders: 1,
					},
				},
			},
		},
		{
			name: "remove a non existing balance",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(0, "EOS"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"lelapinblanc": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"lelapinblanc": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			expectError: true,
		},
		{
			name: "should only change the specific contract",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(0, "EOS"),
					Contract: zsw.AccountName("eidosonecoin"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
				zsw.AccountName("eidosonecoin"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(20, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
				"eidosonecoin": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(30, "EOS"),
								Contract: zsw.AccountName("eidosonecoin"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(20, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
				"eidosonecoin": {},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
				zsw.AccountName("eidosonecoin"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 0,
					},
				},
			},
		},
		{
			name: "remove balance for a non cached contract",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(0, "EOS"),
					Contract: zsw.AccountName("abababababa"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			expectError: true,
		},
		{
			name: "remove balance for non existing token symbol",
			asset: &OwnedAsset{
				Owner: zsw.AccountName("eoscanadadad"),
				Asset: &zsw.ExtendedAsset{
					Asset:    generateTestAsset(0, "WAL"),
					Contract: zsw.AccountName("zswhq.token"),
				},
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			balances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectBalances: map[zsw.AccountName]map[zsw.AccountName][]*OwnedAsset{
				"zswhq.token": {
					"eoscanadadad": {
						{
							Owner: zsw.AccountName("eoscanadadad"),
							Asset: &zsw.ExtendedAsset{
								Asset:    generateTestAsset(100, "EOS"),
								Contract: zsw.AccountName("zswhq.token"),
							},
						},
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Symbol:  "EOS",
						Holders: 1,
					},
				},
			},
			expectError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				TokensInContract: test.tokens,
				Balances:         test.balances,
			}
			err := cache.removeBalance(test.asset)
			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, test.expectBalances, cache.Balances)
			assert.Equal(t, test.expectTokens, cache.TokensInContract)
		})
	}

}

func TestDefaultCache_setToken(t *testing.T) {
	asset := generateTestAsset(1000000, "EOS")
	biggerAsset := generateTestAsset(20000000, "EOS")
	tests := []struct {
		name         string
		token        *pbtokenmeta.Token
		tokens       map[zsw.AccountName][]*pbtokenmeta.Token
		expectTokens map[zsw.AccountName][]*pbtokenmeta.Token
	}{
		{
			name: "sunny path",
			token: &pbtokenmeta.Token{
				Contract:      "zswhq.token",
				Symbol:        "EOS",
				Issuer:        "zswhq.token",
				MaximumSupply: uint64(asset.Amount),
				Precision:     4,
				TotalSupply:   uint64(asset.Amount),
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(asset.Amount),
						Precision:     4,
						TotalSupply:   uint64(asset.Amount),
					},
				},
			},
		},
		{
			name: "update token",
			token: &pbtokenmeta.Token{
				Contract:      "zswhq.token",
				Symbol:        "EOS",
				Issuer:        "zswhq.token",
				MaximumSupply: uint64(biggerAsset.Amount),
				Precision:     4,
				TotalSupply:   uint64(biggerAsset.Amount),
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(asset.Amount),
						Precision:     4,
						TotalSupply:   uint64(asset.Amount),
						Holders:       13,
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(biggerAsset.Amount),
						Precision:     4,
						TotalSupply:   uint64(biggerAsset.Amount),
						Holders:       13,
					},
				},
			},
		},
		{
			name: "add token to existing contract",
			token: &pbtokenmeta.Token{
				Contract:      "zswhq.token",
				Symbol:        "WALL",
				Issuer:        "zswhq.token",
				MaximumSupply: uint64(biggerAsset.Amount),
				Precision:     4,
				TotalSupply:   uint64(biggerAsset.Amount),
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(asset.Amount),
						Precision:     4,
						TotalSupply:   uint64(asset.Amount),
						Holders:       13,
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(asset.Amount),
						Precision:     4,
						TotalSupply:   uint64(asset.Amount),
						Holders:       13,
					},
					{
						Contract:      "zswhq.token",
						Symbol:        "WALL",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(biggerAsset.Amount),
						Precision:     4,
						TotalSupply:   uint64(biggerAsset.Amount),
						Holders:       0,
					},
				},
			},
		},
		{
			name: "add token and contract",
			token: &pbtokenmeta.Token{
				Contract:      "eidosonecoin",
				Symbol:        "EIDOS",
				Issuer:        "eidosonecoin",
				MaximumSupply: uint64(biggerAsset.Amount),
				Precision:     4,
				TotalSupply:   uint64(biggerAsset.Amount),
			},
			tokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(asset.Amount),
						Precision:     4,
						TotalSupply:   uint64(asset.Amount),
						Holders:       7,
					},
				},
			},
			expectTokens: map[zsw.AccountName][]*pbtokenmeta.Token{
				zsw.AccountName("eidosonecoin"): []*pbtokenmeta.Token{
					{
						Contract:      "eidosonecoin",
						Symbol:        "EIDOS",
						Issuer:        "eidosonecoin",
						MaximumSupply: uint64(biggerAsset.Amount),
						Precision:     4,
						TotalSupply:   uint64(biggerAsset.Amount),
						Holders:       0,
					},
				},
				zsw.AccountName("zswhq.token"): []*pbtokenmeta.Token{
					{
						Contract:      "zswhq.token",
						Symbol:        "EOS",
						Issuer:        "zswhq.token",
						MaximumSupply: uint64(asset.Amount),
						Precision:     4,
						TotalSupply:   uint64(asset.Amount),
						Holders:       7,
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				TokensInContract: test.tokens,
			}
			err := cache.setToken(test.token)
			require.NoError(t, err)
			assert.Equal(t, test.expectTokens, cache.TokensInContract)
		})
	}

}

func TestDefaultCache_Stake(t *testing.T) {
	tests := []*struct {
		name           string
		stakeEntries   []*EOSStakeEntry
		expectStakeMap map[zsw.AccountName]*EOSStake
	}{
		{
			name: "golden",
			stakeEntries: []*EOSStakeEntry{
				{
					To:   "b1",
					From: "b1",
					Net:  10000,
					Cpu:  15000,
				},
				{
					To:   "b2",
					From: "b2",
					Net:  20000,
					Cpu:  25000,
				},
				{
					To:   "b3",
					From: "b1",
					Net:  7000,
					Cpu:  13000,
				},
			},
			expectStakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 17000,
					TotalCpu: 28000,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b1"): &EOSStakeEntry{
							To:   "b1",
							From: "b1",
							Net:  10000,
							Cpu:  15000,
						},
						zsw.AccountName("b3"): &EOSStakeEntry{
							To:   "b3",
							From: "b1",
							Net:  7000,
							Cpu:  13000,
						},
					},
				},
				zsw.AccountName("b2"): {
					TotalNet: 20000,
					TotalCpu: 25000,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b2"): &EOSStakeEntry{
							To:   "b2",
							From: "b2",
							Net:  20000,
							Cpu:  25000,
						},
					},
				},
			},
		},
		{
			name: "modify",
			stakeEntries: []*EOSStakeEntry{
				{
					To:   "b1",
					From: "b1",
					Net:  10000,
					Cpu:  15000,
				},
				{
					To:   "b2",
					From: "b2",
					Net:  20000,
					Cpu:  25000,
				},
				{
					To:   "b1",
					From: "b1",
					Net:  0,
					Cpu:  13000,
				},
			},
			expectStakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 0,
					TotalCpu: 13000,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b1"): &EOSStakeEntry{
							To:   "b1",
							From: "b1",
							Net:  0,
							Cpu:  13000,
						},
					},
				},
				zsw.AccountName("b2"): {
					TotalNet: 20000,
					TotalCpu: 25000,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b2"): &EOSStakeEntry{
							To:   "b2",
							From: "b2",
							Net:  20000,
							Cpu:  25000,
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			muts := &MutationsBatch{}
			for _, stakeEntry := range test.stakeEntries {
				muts.SetStake(stakeEntry)
			}
			cache := &DefaultCache{
				EOSStake: make(map[zsw.AccountName]*EOSStake),
			}
			cache.Apply(muts, bstream.NewBlockRef("10a", 10))
			assert.EqualValues(t, test.expectStakeMap, cache.EOSStake)
		})
	}
}

func TestDefaultCache_getStakeForAccount(t *testing.T) {
	tests := []*struct {
		name             string
		account          zsw.AccountName
		stakeMap         map[zsw.AccountName]*EOSStake
		expectStakeValue int64
	}{
		{
			name:    "golden",
			account: zsw.AccountName("b1"),
			stakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 1700,
					TotalCpu: 2800,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b1"): &EOSStakeEntry{
							To:   "b1",
							From: "b1",
							Net:  10000,
							Cpu:  15000,
						},
						zsw.AccountName("b3"): &EOSStakeEntry{
							To:   "b3",
							From: "b1",
							Net:  7000,
							Cpu:  13000,
						},
					},
				},
				zsw.AccountName("b2"): {
					TotalNet: 20000,
					TotalCpu: 25000,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b2"): &EOSStakeEntry{
							To:   "b2",
							From: "b2",
							Net:  20000,
							Cpu:  25000,
						},
					},
				},
			},
			expectStakeValue: 4500,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				EOSStake: test.stakeMap,
			}
			assert.EqualValues(t, test.expectStakeValue, cache.getStakeForAccount(test.account))
		})
	}
}

func TestDefaultCache_setStake(t *testing.T) {
	tests := []*struct {
		name           string
		stakeEntry     *EOSStakeEntry
		stakeMap       map[zsw.AccountName]*EOSStake
		expectStakeMap map[zsw.AccountName]*EOSStake
	}{
		{
			name: "no stake entry present",
			stakeEntry: &EOSStakeEntry{
				To:   "b2",
				From: "b1",
				Net:  1200,
				Cpu:  2400,
			},
			stakeMap: map[zsw.AccountName]*EOSStake{},
			expectStakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 1200,
					TotalCpu: 2400,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b2"): &EOSStakeEntry{
							To:   "b2",
							From: "b1",
							Net:  1200,
							Cpu:  2400,
						},
					},
				},
			},
		},
		{
			name: "stake entry present with a different stake.to",
			stakeEntry: &EOSStakeEntry{
				To:   "b1",
				From: "b1",
				Net:  70000,
				Cpu:  140000,
			},
			stakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 1200,
					TotalCpu: 2400,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b2"): &EOSStakeEntry{
							To:   "b2",
							From: "b1",
							Net:  1200,
							Cpu:  2400,
						},
					},
				},
			},
			expectStakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 71200,
					TotalCpu: 142400,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b2"): &EOSStakeEntry{
							To:   "b2",
							From: "b1",
							Net:  1200,
							Cpu:  2400,
						},
						zsw.AccountName("b1"): &EOSStakeEntry{
							To:   "b1",
							From: "b1",
							Net:  70000,
							Cpu:  140000,
						},
					},
				},
			},
		},
		{
			name: "stake entry present with an already exisiting stake.to",
			stakeEntry: &EOSStakeEntry{
				To:   "b1",
				From: "b1",
				Net:  70000,
				Cpu:  140000,
			},
			stakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 1200,
					TotalCpu: 2400,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b1"): &EOSStakeEntry{
							To:   "b1",
							From: "b1",
							Net:  1200,
							Cpu:  2400,
						},
					},
				},
			},
			expectStakeMap: map[zsw.AccountName]*EOSStake{
				zsw.AccountName("b1"): {
					TotalNet: 70000,
					TotalCpu: 140000,
					Entries: map[zsw.AccountName]*EOSStakeEntry{
						zsw.AccountName("b1"): &EOSStakeEntry{
							To:   "b1",
							From: "b1",
							Net:  70000,
							Cpu:  140000,
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &DefaultCache{
				EOSStake: test.stakeMap,
			}
			cache.setStake(test.stakeEntry)
			assert.EqualValues(t, test.expectStakeMap, cache.EOSStake)
		})
	}
}

func generateTestAsset(amount zsw.Int64, symbol string) zsw.Asset {
	return zsw.Asset{
		Amount: amount,
		Symbol: *generateTestSymbol(symbol),
	}
}

func generateTestSymbol(symbol string) *zsw.Symbol {
	return &zsw.Symbol{
		Precision: 4,
		Symbol:    symbol,
	}
}
