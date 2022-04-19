package monitors

import (
	"fmt"
	"os/exec"
	"strings"
)

type PingMonitor struct {
	Host string
}

func (m *PingMonitor) Monitor() bool {
	fmt.Printf("Start pinging host '%s'\n", m.Host)
	out, _ := exec.Command("ping", m.Host, "-c 5", "-i 3", "-w 10").Output()
	return !strings.Contains(string(out), "Destination Host Unreachable")
}
