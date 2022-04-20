package common

import (
	"fmt"
	"gonitor/monitors"
	"os"
	"strings"
	"time"
)

func Cli(args Arguments) {
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
			m := monitors.PingMonitor{Host: host}
			fmt.Printf("%v -> ", host)
			ok, _ := m.Monitor()
			res = ok
			break
		case "rest":
			url := args.GetFirstRequired("u", "url")
			m := monitors.RestMonitor{URL: url}
			fmt.Printf("%v -> ", url)
			ok, _ := m.Monitor()
			res = ok
			break
		case "ssh":
			host := args.GetFirstRequired("H", "host")
			port := args.GetFirstDefault("22", "p", "port")
			user := args.GetFirstRequired("u", "user")
			pass := args.GetFirstRequired("P", "password", "pass")
			m := monitors.SSHMonitor{
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
			ok, _ := m.Monitor()
			res = ok
			break
		case "port":
			host := args.GetFirstRequired("H", "host")
			port := args.GetFirstRequired("p", "port")
			m := monitors.PortMonitor{
				Host: host,
				Port: port,
			}
			fmt.Printf("%v:%v -> ", host, port)
			ok, _ := m.Monitor()
			res = ok
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
