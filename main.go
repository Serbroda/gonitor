package main

import (
	"fmt"
	"gonitor/common"
)

func main() {
	args := common.GetArgs()

	if args.HasAnyKey("m", "mode") {
		common.Cli(args)
		return
	}

	configFile := args.GetFirstDefault("test.yml", "c", "config")
	conf := common.LoadConfig(configFile)
	fmt.Printf("Conf: %v\n", conf)
}
