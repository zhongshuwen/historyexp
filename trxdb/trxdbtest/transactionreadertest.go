// Copyright 2019 dfuse Platform Inc.
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

package trxdbtest

import (
	"context"
	"fmt"
	"testing"

	ct "github.com/dfuse-io/dfuse-eosio/codec/testing"
	pbcodec "github.com/dfuse-io/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	"github.com/dfuse-io/dfuse-eosio/trxdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var transactionReaderTests = []DriverTestFunc{
	TestGetTransactionTraces,
	TestGetTransactionTracesBatch,
	TestGetTransactionEvents,
	TestGetTransactionEventsBatch,
	TestReadTransactions,
}

func TestReadTransactions(t *testing.T, driverFactory DriverFactory) {
	db, clean := driverFactory()
	defer clean()

	ctx := context.Background()
	in := testBlock1()

	require.NoError(t, db.PutBlock(ctx, in))
	require.NoError(t, db.UpdateNowIrreversibleBlock(ctx, in))
	require.NoError(t, db.Flush(ctx))

	// Block data
	evs, err := db.GetTransactionEvents(context.Background(), "00112233aa")
	require.NoError(t, err)
	assert.Len(t, evs, 2)
	var additions, executions int

	for _, ev := range evs {
		switch evt := ev.Event.(type) {
		case *pbcodec.TransactionEvent_Addition:
			additions++
			assert.Equal(t, "00112233aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ev.Id)
			assert.Equal(t, "00000002aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ev.BlockId)
			assert.True(t, ev.Irreversible)

			if evt.Addition.Receipt != nil {
				// FIXME: this should be skipped ONLY for the old `bigt` implementation, which didn't save
				// the Receipt..
				assert.Equal(t, 32, int(evt.Addition.Receipt.NetUsageWords))
				assert.Equal(t, 32, int(evt.Addition.Receipt.CpuUsageMicroSeconds))
			}
			assert.Equal(t, []string{"SIG_K1_K7kTcvsznS2pSQ2unjW9nduqHieWnc5B6rFdbVif4RM1DCTVhQUpzwng3XTGewDhVZqNvqSAEwHgB8yBnfDYAHquRX4fBo"}, evt.Addition.Transaction.Signatures)
			assert.Len(t, evt.Addition.Transaction.Transaction.Actions, 1)
			assert.Equal(t, "name", evt.Addition.Transaction.Transaction.Actions[0].Name)
			assert.Equal(t, []string{"EOS7T3GcBYpYf2D63HGDG7qB9TiD56XT4m1hAQfkHWuV9LhMoQ1ZY"}, evt.Addition.PublicKeys.PublicKeys)

		case *pbcodec.TransactionEvent_Execution:
			executions++
			assert.Equal(t, "00112233aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ev.Id)
			assert.Equal(t, "00000002aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ev.BlockId)
			assert.True(t, ev.Irreversible)

			assert.Equal(t, "00000001aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", evt.Execution.BlockHeader.Previous)
			assert.Equal(t, "tester", evt.Execution.BlockHeader.Producer)
			assert.Len(t, evt.Execution.Trace.DtrxOps, 2)
			assert.Equal(t, "aaa888aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", evt.Execution.Trace.DtrxOps[1].TransactionId)
			assert.Equal(t, "00112233aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", evt.Execution.Trace.Id)

		default:
			t.Error(fmt.Sprintf("unexpected type %T", ev))
		}
	}

	assert.Equal(t, 1, additions)
	assert.Equal(t, 1, executions)
}

func TestGetTransactionTraces(t *testing.T, driverFactory DriverFactory) {
	tests := []struct {
		name        string
		trxIDs      []string
		trxIDPrefix string
		expectTrxID string
		expectErr   error
	}{
		{
			name:        "sunny path",
			trxIDs:      []string{"a1bc5790ef36d5779e2a0a849a11c09c999b5dc564afce6920e20b07af1f4b6a", "a2bc5790ef36d5779e2a0a849a11c09c999b5dc564afce6920e20b07af1f4b6a", "a3bc5790ef36d5779e2a0a849a11c09c999b5dc564afce6920e20b07af1f4b6b"},
			trxIDPrefix: "a1",
			expectTrxID: "a1bc5790ef36d5779e2a0a849a11c09c999b5dc564afce6920e20b07af1f4b6a",
		},
		{
			name:        "multiple matches error",
			trxIDs:      []string{"a1bc000000000000000000000000000000000000000000000000000000000000", "a1bc111111111111111111111111111111111111111111111111111111111111", "a1bc222222222222222222222222222222222222222222222222222222222222"},
			trxIDPrefix: "a1",
			expectErr:   fmt.Errorf("requested prefix a1 returns multiple transactions (a1bc000000000000000000000000000000000000000000000000000000000000, a1bc111111111111111111111111111111111111111111111111111111111111...)"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, trxID := range test.trxIDs {
				putTransaction(t, db, trxID)
			}

			events, err := db.GetTransactionTraces(ctx, test.trxIDPrefix)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				require.Len(t, events, 1)
				assert.Equal(t, test.expectTrxID, events[0].Id)
			}
		})
	}
}

func TestGetTransactionTracesBatch(t *testing.T, driverFactory DriverFactory) {
	tests := []struct {
		name         string
		trxIDs       []string
		trxIdsPrefix []string
		expectTrxIDs []string
		expectErr    error
	}{
		{
			name:         "sunny path",
			trxIDs:       []string{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "2abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "3aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "3abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
			trxIdsPrefix: []string{"1aaa", "1abb", "2aaa", "2abb"},
			expectTrxIDs: []string{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "2abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
		},
		{
			name:         "multiple matches error",
			trxIDs:       []string{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "2abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "3aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "3abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
			trxIdsPrefix: []string{"1a", "2a"},
			expectErr:    fmt.Errorf("requested prefix 1a returns multiple transactions (1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa, 1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb...)"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ctx = context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, trxID := range test.trxIDs {
				putTransaction(t, db, trxID)
			}

			matches, err := db.GetTransactionTracesBatch(ctx, test.trxIdsPrefix)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				eventIds := []string{}
				for _, trxEvents := range matches {
					require.Len(t, trxEvents, 1)
					eventIds = append(eventIds, trxEvents[0].Id)
				}
				assert.ElementsMatch(t, test.expectTrxIDs, eventIds)
			}
		})
	}
}

func TestGetTransactionEvents(t *testing.T, driverFactory DriverFactory) {
	tests := []struct {
		name        string
		trxIDs      []string
		trxIDPrefix string
		expectTrxID string
		expectErr   error
	}{
		{
			name:        "sunny path",
			trxIDs:      []string{"1abbffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "1accffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "1eddffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"},
			trxIDPrefix: "1abbff",
			expectTrxID: "1abbffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		},
		{
			name:        "full match",
			trxIDs:      []string{"1abbffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "1accffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "1eddffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"},
			trxIDPrefix: "1eddffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			expectTrxID: "1eddffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ctx = context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, trxID := range test.trxIDs {
				putTransaction(t, db, trxID)
			}

			// TODO: the `GetTransactionEvents()` function should be
			// exercised with all the types of events. So fixtures
			// should write an implicit trx, two addition events
			// (internal and normal), a dtrx event and an execution
			// trace event.

			events, err := db.GetTransactionEvents(ctx, test.trxIDPrefix)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				require.Len(t, events, 2)
				require.Equal(t, events[0].Id, test.expectTrxID)
				require.Equal(t, events[1].Id, test.expectTrxID)
			}

		})
	}
}

func TestGetTransactionEventsBatch(t *testing.T, driverFactory DriverFactory) {
	tests := []struct {
		name         string
		trxIDs       []string
		trxIdsPrefix []string
		expectTrxIDs []string
		expectErr    error
	}{
		{
			name:         "sunny path",
			trxIDs:       []string{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "2abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "3aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "3abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
			trxIdsPrefix: []string{"1aaa", "1abb", "2aaa", "2abb"},
			expectTrxIDs: []string{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "2abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
		},
		{
			name:         "multiple matches error",
			trxIDs:       []string{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "2abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "3aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "3abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
			trxIdsPrefix: []string{"1a", "2a"},
			expectErr:    fmt.Errorf("requested prefix 1a returns multiple transactions (1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa, 1abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb...)"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ctx = context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, trxID := range test.trxIDs {
				putTransaction(t, db, trxID)
			}

			matches, err := db.GetTransactionEventsBatch(ctx, test.trxIdsPrefix)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				eventIds := []string{}
				for _, trxEvents := range matches {
					require.Len(t, trxEvents, 2)
					require.Equal(t, trxEvents[0].Id, trxEvents[1].Id)
					eventIds = append(eventIds, trxEvents[0].Id)
				}
				assert.ElementsMatch(t, test.expectTrxIDs, eventIds)
			}
		})
	}
}

func putTransaction(t *testing.T, db trxdb.DB, trxID string) {
	// Need to use a full block id string (64 characters, 32 bytes) because keys transaction trace key unpacking
	// expects a full length block id, you get `invalid key length` errors if not long enough
	blk := ct.Block(t, "00000002aa000000000000000000000000000000000000000000000000000000",
		ct.TrxTrace(t, ct.TrxID(trxID),
			// FIXME: a dtrx that is created actually has a *different* transaction ID from the one creating it.
			ct.DtrxOp(t, "create", trxID, ct.DtrxOpPayer("eoscanada1"), &pbcodec.SignedTransaction{
				Transaction:     nil,
				Signatures:      []string{"signature"},
				ContextFreeData: nil,
			}),
		),
	)

	ctx := context.Background()
	require.NoError(t, db.PutBlock(ctx, blk))
	require.NoError(t, db.UpdateNowIrreversibleBlock(ctx, blk))
	require.NoError(t, db.Flush(ctx))
}
