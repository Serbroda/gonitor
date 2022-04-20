package common

import (
	"fmt"
	"gonitor/monitors"
	"os"
	"time"
)

func Cli(args Arguments) {
	mode := args.GetFirstRequired("m", "mode")
	repeat := false
	if args.HasAnyKey("r", "repeat") {
		repeat = true
	}
	res := false

	m := monitors.NewMonitor(monitors.MonitorType(mode), args.keyValues)
	for {
		ok, _ := m.Monitor()
		res = ok
		fmt.Printf("%v\n", res)
		if !repeat {
			break
		}
		res = false
		time.Sleep(5 * time.Second)
	}

	if res {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
