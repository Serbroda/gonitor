package main

import (
	"fmt"
	"gonitor/common"
	"gonitor/monitors"
	"os"
	"strings"
	"time"
)

func main() {
	args := common.GetArgs()
	mode := args.GetFirstRequired("m", "mode")
	repeat := false
	if args.HasAnyKey("r", "repeat") {
		repeat = true
	}
	res := false

	for {
		switch strings.ToLower(mode) {
		case "ping":
			host := args.GetFirstRequired("H", "host")
			ping := monitors.PingMonitor{Host: host}
			fmt.Printf("%v -> ", host)
			res = ping.Monitor()
			break
		case "rest":
			url := args.GetFirstRequired("u", "url")
			rest := monitors.RestMonitor{URL: url}
			fmt.Printf("%v -> ", url)
			res = rest.Monitor()
			break
		case "ssh":
			host := args.GetFirstRequired("H", "host")
			port := args.GetFirstDefault("22", "p", "port")
			user := args.GetFirstRequired("u", "user")
			pass := args.GetFirstRequired("P", "password", "pass")
			ssh := monitors.SSHMonitor{
				Host:     host,
				Port:     port,
				User:     user,
				Password: pass,
				Command:  "df -h /",
				ResultParser: func(out string) bool {
					return true
				},
			}
			fmt.Printf("%v -> ", host)
			res = ssh.Monitor()
			break
		default:
			panic("Unkown mode: " + mode)
		}

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
