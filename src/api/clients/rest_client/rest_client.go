package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RestClient allows assignment of any type that implements *http.Client's do func
type RestClient interface {
	Do(*http.Request) (*http.Response, error)
}

var (
	Client RestClient
)

func init() {
	Client = &http.Client{}
}

// POST makes an HTTP Post request and returns an HTTP response
func Post(headers http.Header, url string, data interface{}) (*http.Response, error) {

	js, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(js))
	if err != nil {
		return nil, fmt.Errorf("could not create a new http request, %v", err)
	}

	request.Header = headers

	return Client.Do(request)
}
