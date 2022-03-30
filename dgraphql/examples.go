package dgraphql

import (
	"encoding/json"
	"fmt"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/invisible-train-40/dgraphql/static"
)

//go:generate rice embed-go

// GraphqlExamples returns the ordered list of predefined GraphQL examples that should be
// displayed inside GraphiQL interface.
func GraphqlExamples(config *Config) []*static.GraphqlExample {
	box := rice.MustFindBox("graphql-examples")
	examples := []*static.GraphqlExample{
		{
			Label:    "搜索信息流",
			Document: graphqlDocument(box, "search_stream_forward.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"query": "receiver:zsw.items action:mint", "cursor": "", "limit": 100}`),
			},
		},
		{
			Label:    "搜索信息流 (反向排序)",
			Document: graphqlDocument(box, "search_stream_backward.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"query": "receiver:zsw.items action:transfer", "cursor": "", "low": 1, "limit": 10}`),
			},
		},
		{
			Label:    "搜索查询",
			Document: graphqlDocument(box, "search_query_forward.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"query": "receiver:zswhq action:newaccount", "cursor": "", "limit": 10}`),
			},
		},
		{
			Label:    "搜索查询（反向排序）",
			Document: graphqlDocument(box, "search_query_backward.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"query": "receiver:zswhq action:newaccount", "low": -500, "high": -1, "cursor": "", "limit": 10}`),
			},
		},
		{
			Label:    "时间范围查询",
			Document: graphqlDocument(box, "time_ranges.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(fmt.Sprintf(`{"start": "%s", "end": "%s"}`, dateOffsetByBlock(0), dateOffsetByBlock(5))),
				"mainnet": r(fmt.Sprintf(`{"start": "%s", "end": "%s"}`, oneWeekAgo(), dateOffsetByBlock(-1))),
				"jungle":  r("mainnet"),
				"kylin":   r("mainnet"),
			},
		},
		{
			Label:    "查询区块信息（Hash ID）",
			Document: graphqlDocument(box, "get_block_by_id.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"blockId": "<Block ID Here>"}`),
				"mainnet": r(`{"blockId": "063a7e525142f64d7465bbebc690afbb228bff7d7e0ffda31d9a06106fbc1982"}`),
				"jungle":  r(`{"blockId": "047c7822f396e64b9cbb28cc2b199b8e5a4c33c894b8742eab646e670486bb0d"}`),
				"kylin":   r(`{"blockId": "05609e94b57cdea5ce4ff8afa89070d37a85923855f0d41efdbd956dbaddb5f7"}`),
			},
		},
		{
			Label:    "查询区块信息（n-th）",
			Document: graphqlDocument(box, "get_block_by_num.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"blockNum": 10}`),
			},
		},
		{
			Label:    "查询剩下的计算资源",
			Document: graphqlDocument(box, "get_account_balances.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"account": "zsw.admin", "opts": ["ZSWCC_INCLUDE_STAKED"], "limit": 10}`),
			},
		},
	}

	if config.AccountHistAccountAddr != "" {
		examples = append(examples, &static.GraphqlExample{
			Label:    "查询账号历史",
			Document: graphqlDocument(box, "get_account_history_by_account.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"account": "zsw.admin","limit": 100}`),
				"mainnet": r(`{"account": "zsw.admin","limit": 100}`),
			},
		})
	}

	if config.AccountHistAccountContractAddr != "" {
		examples = append(examples, &static.GraphqlExample{
			Label:    "查询账号历史（智能合约）",
			Document: graphqlDocument(box, "get_account_history_by_account_contract.graphql"),
			Variables: static.GraphqlVariablesByNetwork{
				"generic": r(`{"account": "zswhq", "contract": "zswhq.token", "limit": 100}`),
				"mainnet": r(`{"account": "eoscanadacom", "contract": "zswhq.token", "limit": 100}`),
				"kylin":   r("mainnet"),
				"dev1":    r(`{"account": "battlefield1", "contract": "zswhq.token", "limit": 100}`),
			},
		})
	}

	return examples
}

func graphqlDocument(box *rice.Box, name string) static.GraphqlDocument {
	asset, err := box.String(name)
	if err != nil {
		panic(fmt.Errorf("unable to get content for graphql examples file %q: %w", name, err))
	}

	return static.GraphqlDocument(asset)
}

func oneWeekAgo() string {
	return time.Now().Add(-7 * 24 * time.Hour).UTC().Format("2006-01-02T15:04:05Z")
}

func dateOffsetByBlock(blockCount int) string {
	return time.Now().Add(time.Duration(blockCount) * 500 * time.Millisecond).UTC().Format("2006-01-02T15:04:05Z")
}

func r(rawJSON string) json.RawMessage {
	return json.RawMessage(rawJSON)
}
