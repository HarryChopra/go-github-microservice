package mocks

import (
	"io"
	"net/http"
)

type MockClientImpl struct{}

var (
	do func(*http.Request) (*http.Response, error)
)

func (c *MockClientImpl) Do(r *http.Request) (*http.Response, error) {
	return do(r)
}

// SetDoFunc sets a custom behaviour for the mocked http.Client's Do func
func SetDoFunc(body io.ReadCloser, statusCode int, err error) {
	var response *http.Response

	if body == nil {
		response = nil
	} else {
		response = &http.Response{
			StatusCode: statusCode,
			Body:       body,
		}
	}

	do = func(*http.Request) (*http.Response, error) {
		return response, err
	}
}
