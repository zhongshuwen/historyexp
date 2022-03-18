package blockmeta

import (
	"context"

	"github.com/dfuse-io/blockmeta"
	"github.com/zhongshuwen/zswchain-go"
)

func init() {
	blockmeta.GetBlockNumFromID = blockNumFromID
}

func blockNumFromID(ctx context.Context, id string) (uint64, error) {
	return uint64(eos.BlockNum(id)), nil
}
