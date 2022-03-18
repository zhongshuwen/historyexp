package main

import (
	"github.com/zhongshuwen/historyexp/cmd/dfuseeos/cli"
)

var version = "dev"
var commit = ""

func init() {
	cli.RootCmd.Version = version + "-" + commit
}

func main() {
	cli.Main()
}
