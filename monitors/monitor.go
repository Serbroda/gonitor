package monitors

import (
	"gonitor/utils"
)

type MonitorType string

const (
	Ping MonitorType = "ping"
	SSH              = "ssh"
	REST             = "rest"
	Port             = "port"
)

type Monitor interface {
	Monitor() (bool, any)
}

func NewMonitor(monitorType MonitorType, properties map[string]string) Monitor {
	switch monitorType {
	case Ping:
		host := utils.GetFirstRequired(properties, "H", "host")
		return &PingMonitor{Host: host}
	case REST:
		url := utils.GetFirstRequired(properties, "u", "url")
		return &RestMonitor{URL: url}
	case SSH:
		host := utils.GetFirstRequired(properties, "H", "host")
		port := utils.GetFirstDefault(properties, "22", "p", "port")
		user := utils.GetFirstRequired(properties, "u", "user")
		pass := utils.GetFirstRequired(properties, "P", "password", "pass")
		return &SSHMonitor{
			Host:     host,
			Port:     port,
			User:     user,
			Password: pass,
			Command:  "df -h /",
			ResultParser: func(out string) bool {
				return true
			},
		}
	case Port:
		host := utils.GetFirstRequired(properties, "H", "host")
		port := utils.GetFirstRequired(properties, "p", "port")
		return &PortMonitor{
			Host: host,
			Port: port,
		}
		break
	default:
		panic("Unkown monitor type: " + monitorType)
	}
	return nil
}

//func NewMonitors(confs []common.MonitorConfig) []Monitor {
//	monitors := make([]Monitor, len(confs))
//
//	for _, mf := range confs {
//		m := NewMonitor(mf.Type, mf.Properties)
//		monitors = append(monitors, m)
//	}
//	return monitors
//}
