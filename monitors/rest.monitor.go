package monitors

import (
	"fmt"
	"net/http"
)

type RestMonitor struct {
	URL string
}

func (m *RestMonitor) Monitor() bool {
	fmt.Printf("Start calling rest url '%s'\n", m.URL)
	response, err := http.Get(m.URL)
	if err != nil {
		fmt.Printf("Failed to call rest URL %v\n", err)
		return false
	}
	return response.StatusCode > 99 && response.StatusCode < 300
}
