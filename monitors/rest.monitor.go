package monitors

import (
	"net/http"
)

type RestMonitor struct {
	URL string
}

func (m *RestMonitor) Monitor() bool {
	response, err := http.Get(m.URL)
	if err != nil {
		return false
	}
	return response.StatusCode > 99 && response.StatusCode < 300
}
