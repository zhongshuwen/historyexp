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

package eosws

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/dfuse-io/bstream"
	"github.com/zhongshuwen/historyexp/codec"
	pbcodec "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/codec/v1"
	"github.com/dfuse-io/jsonpb"
	"github.com/dfuse-io/logging"
zsw "github.com/zhongshuwen/zswchain-go"
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func init() {
	logging.TestingOverride()
}

func testBlock(t *testing.T, id, previousID, producer string, libNum uint64, trxTraceJSONs ...string) *bstream.Block {
	trxTraces := make([]*pbcodec.TransactionTrace, len(trxTraceJSONs))
	for i, trxTraceJSON := range trxTraceJSONs {
		trxTrace := new(pbcodec.TransactionTrace)

		require.NoError(t, jsonpb.UnmarshalString(trxTraceJSON, trxTrace))

		trxTraces[i] = trxTrace
	}

	pbblock := &pbcodec.Block{
		Id:     id,
		Number: zsw.BlockNum(id),
		Header: &pbcodec.BlockHeader{
			Previous:  previousID,
			Producer:  producer,
			Timestamp: &timestamp.Timestamp{},
		},
		UnfilteredTransactionTraces: trxTraces,
	}

	if os.Getenv("DEBUG") != "" {
		out, err := json.Marshal(pbblock)
		require.NoError(t, err)

		// We re-normalize to a plain map[string]interface{} so it's printed as JSON and not a proto default String implementation
		normalizedOut := map[string]interface{}{}
		require.NoError(t, json.Unmarshal(out, &normalizedOut))

		zlog.Debug("created test block", zap.Any("block", normalizedOut))
	}

	block, err := codec.BlockFromProto(pbblock)
	require.NoError(t, err)

	return block
}

func protoJSONMarshalIndent(t *testing.T, message proto.Message) string {
	value, err := jsonpb.MarshalIndentToString(message, "  ")
	require.NoError(t, err)

	return value
}
