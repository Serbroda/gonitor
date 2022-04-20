package monitors

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type RestMonitor struct {
	URL string
}

type RestResponse struct {
	StatusCode int
	Body       string
}

func (m *RestMonitor) Monitor() (bool, any) {
	response, err := http.Get(m.URL)
	if err != nil {
		return false, RestResponse{}
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response")
		panic(err)
	}
	ok := response.StatusCode > 99 && response.StatusCode < 300
	return ok, RestResponse{
		StatusCode: response.StatusCode,
		Body:       string(responseData),
	}
}
