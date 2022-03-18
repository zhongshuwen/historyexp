package grpc

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func init() {
	logging.Register("github.com/zhongshuwen/historyexp/accounthist/grpc", &zlog)
}
