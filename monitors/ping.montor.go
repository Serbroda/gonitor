package monitors

import (
	"github.com/go-ping/ping"
	"time"
)

type PingMonitor struct {
	Host string
}

func (m *PingMonitor) Monitor() (bool, any) {
	pinger, err := ping.NewPinger(m.Host)
	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Timeout = 5 * time.Second

	if err != nil {
		panic(err)
	}
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	return stats.PacketsRecv > 0, nil
}
