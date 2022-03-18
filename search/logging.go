package search

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var traceEnabled = logging.IsTraceEnabled("search", "github.com/zhongshuwen/historyexp/search")
var zlog *zap.Logger

func init() {
	logging.Register("github.com/zhongshuwen/historyexp/search", &zlog)
}
