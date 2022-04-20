package monitors

import (
	"fmt"
	"net"
	"time"
)

type PortMonitor struct {
	Host string
	Port string
}

func (m *PortMonitor) Monitor() (bool, any) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(m.Host, m.Port), 5*time.Second)
	if err != nil {
		fmt.Println("Connection error: ", err)
		return false, nil
	}
	defer conn.Close()
	return conn != nil, nil
}
