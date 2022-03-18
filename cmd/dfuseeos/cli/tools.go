package cli

import "github.com/zhongshuwen/historyexp/tools"

func init() {
	RootCmd.AddCommand(tools.Cmd)
}
