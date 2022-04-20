package main

import (
	"gonitor/common"
)

func main() {
	args := common.GetArgs()

	if args.HasAnyKey("m", "mode") {
		common.Cli(args)
		return
	}
}
