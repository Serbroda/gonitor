package main

import (
	"gonitor/common"
	"gonitor/tui"
)

func main() {
	args := common.GetArgs()

	if args.HasAnyKey("m", "mode") {
		common.Cli(args)
		return
	}

	//configFile := args.GetFirstDefault("test.yml", "c", "config")
	//conf := common.LoadConfig(configFile)
	//fmt.Printf("Conf: %v\n", conf)

	/*mons := []monitors.Monitor{
		monitors.NewMonitor(monitors.REST, map[string]string{"url": "http://www.google.de"}),
		monitors.NewMonitor(monitors.REST, map[string]string{"url": "http://www.fds.de"}),
		monitors.NewMonitor(monitors.REST, map[string]string{"url": "http://www.heise.de"}),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		i := 0
		for {
			for _, m := range mons {
				ok, _ := m.Monitor()
				fmt.Printf("Res: %v\n", ok)
			}
			time.Sleep(2 * time.Second)
			i++
			if i > 10 {
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()*/
	tui.Start()
}
