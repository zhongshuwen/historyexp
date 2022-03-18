package filtering

import (
	"testing"

	ct "github.com/dfuse-io/dfuse-eosio/codec/testing"
	pbcodec "github.com/dfuse-io/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlockFilter(t *testing.T) {
	filterMatched := true
	filterDidNotMatch := false

	tests := []struct {
		name           string
		exprs          filters
		trace          *pbcodec.TransactionTrace
		expectedPass   bool
		expectedSystem bool
	}{
		{
			"filter nothing",
			getFilters("", "", ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "whatever:action")),
			filterMatched, false,
		},
		{
			"filter nothing, with default programs",
			getFilters("true", "false", ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "whatever:action")),
			filterMatched, false,
		},
		{
			"blacklist things FROM badguy",
			getFilters(`true`, `account == "zswhq.token" && data.from == "badguy"`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq.token:transfer", ct.ActionData(`{"from":"goodguy","to":"badguy"}`))),
			filterMatched, false,
		},
		{
			"blacklist things TO badguy",
			getFilters(`true`, "account == 'zswhq.token' && data.to == 'badguy'", ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq.token:transfer", ct.ActionData(`{"from":"goodguy","to":"badguy"}`))),
			filterDidNotMatch, false,
		},
		{
			"blacklist transfers to eidosonecoin",
			getFilters(
				"*",
				`account == 'eidosonecoin' || receiver == 'eidosonecoin' || (account == 'zswhq.token' && (data.to == 'eidosonecoin' || data.from == 'eidosonecoin'))`,
				"",
			),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq.token:transfer", ct.ActionData(`{"from":"goodguy","to":"eidosonecoin"}`))),
			filterDidNotMatch, false,
		},
		{
			"non-matching identifier in exclude-filter program doesn't blacklist",
			getFilters("", `account == 'zswhq.token' && data.from == 'broken'`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq.token:issue", ct.ActionData(`{"to":"winner"}`))),
			filterMatched, false,
		},
		{
			"a key not found error in include-filter still includes transaction",
			getFilters(`account == 'zswhq.token' && data.bob == 'broken'`, "", ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq.token:issue", ct.ActionData(`{"to":"winner"}`))),
			filterMatched, false,
		},
		{
			"both whitelist and blacklist fail",
			getFilters(`data.bob == 'broken'`, `data.rita == 'rebroken'`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "any:any", ct.ActionData(`{"denise":"winner"}`))),
			filterMatched, false,
		},
		{
			"whitelisted but blacklist cleans out",
			getFilters(`data.bob == '1'`, `data.rita == '2'`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "any:any", ct.ActionData(`{"bob":"1","rita":"2"}`))),
			false, false,
		},
		{
			"whitelisted but blacklist broken so doesn't clean out",
			getFilters(`data.bob == '1'`, `data.broken == 'really'`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "any:any", ct.ActionData(`{"bob":"1"}`))),
			filterMatched, false,
		},

		{
			"block receiver",
			getFilters("", `receiver == "badguy"`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "badguy:any:any", ct.ActionData(`{}`))),
			filterDidNotMatch, false,
		},
		{
			"block spam",
			getFilters("", `(trx_action_count > 20 && top5_trx_actors.exists(x, x in ['badguy','worstguy']))`, ""),
			ct.TrxTrace(t, multiActionTraces(t, "badguy:any:any", 30, ct.ActionData(`{}`))...),
			filterDidNotMatch, false,
		},
		{
			"block only spam",
			getFilters("", `(trx_action_count > 20 && top5_trx_actors.exists(x, x in ['badguy','worstguy']))`, ""),
			ct.TrxTrace(t, multiActionTraces(t, "badguy:any:any", 10, ct.ActionData(`{}`))...),
			filterMatched, false,
		},
		{
			"block only specific",
			getFilters("", `(trx_action_count > 20 && top5_trx_actors.exists(x, x in ['badguy','worstguy']))`, ""),
			ct.TrxTrace(t, multiActionTraces(t, "goodguy:any:any", 30, ct.ActionData(`{}`))...),
			filterMatched, false,
		},
		{
			"prevent a failure on evaluation, so matches because blacklist fails",
			getFilters("", `account == "badacct" && has(data.from) && data.from != "badguy"`, ""),
			ct.TrxTrace(t, ct.ActionTrace(t, "badrecv:badacct:any", ct.ActionData(`{}`))),
			filterMatched, false,
		},

		{
			"system action already included are not flagged as system",
			getFilters(`action == "setabi"`, ``, `action == "setabi"`),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq:zswhq:setabi", ct.ActionData(`{}`))),
			filterMatched, false,
		},
		{
			"system action are included even when not included",
			getFilters(`action == "anythingelse"`, ``, `action == "setabi"`),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq:zswhq:setabi", ct.ActionData(`{}`))),
			filterMatched, true,
		},
		{
			"system action are included even when excluded",
			getFilters("*", `action == "setabi"`, `action == "setabi"`),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq:zswhq:setabi", ct.ActionData(`{}`))),
			filterMatched, true,
		},
		{
			"system action are included even when excluded and not included",
			getFilters(`action == "anythingelse"`, `action == "setabi"`, `action == "setabi"`),
			ct.TrxTrace(t, ct.ActionTrace(t, "zswhq:zswhq:setabi", ct.ActionData(`{}`))),
			filterMatched, true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.GreaterOrEqual(t, len(test.trace.ActionTraces), 1, "This test requires action traces in trxTraces")

			filter, err := NewBlockFilter(test.exprs.include, test.exprs.exclude, test.exprs.system)
			require.NoError(t, err)

			hasPass, isSystem := shouldProcess(
				&MemoizableTrxTrace{TrxTrace: test.trace},
				test.trace.ActionTraces[0],
				filter.IncludeProgram.choose(0),
				filter.ExcludeProgram.choose(0),
				filter.SystemActionsIncludeProgram.choose(0),
			)
			if test.expectedPass {
				assert.True(t, hasPass, "Expected action trace to match filter (%s) but it did not", test.exprs)
			} else {
				assert.False(t, hasPass, "Expected action trace to NOT match filter (%s) but it did", test.exprs)
			}

			if test.expectedSystem {
				assert.True(t, isSystem, "Expected action trace to be system included (%s) but it did not", test.exprs)
			} else {
				assert.False(t, isSystem, "Expected action trace to NOT be system included (%s) but it did", test.exprs)
			}
		})
	}
}

func getFilters(a string, b string, c string) filters {
	return filters{
		[]string{a},
		[]string{b},
		[]string{c},
	}
}

type filters struct {
	include []string
	exclude []string
	system  []string
}

func multiActionTraces(t *testing.T, receiverAccountActionNameTriplet string, repeats int, components ...interface{}) (out []interface{}) {
	for i := 0; i < repeats; i++ {
		out = append(out, ct.ActionTrace(t, receiverAccountActionNameTriplet, components...))
	}
	return
}
