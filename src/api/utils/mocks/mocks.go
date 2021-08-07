package mocks

import (
	"fmt"
	"io"
	"net/http"
)

// MockClient extends RestClient interface
type MockClient interface {
	Do(*http.Request) (*http.Response, error)
	SetDoFunc(io.ReadCloser, int, error)
}

// MockClient implementation
type MockClientImpl struct {
	DoFunc func(r *http.Request) (*http.Response, error)
}

// SetDoFunc sets the do function for mock client
func (c *MockClientImpl) SetDoFunc(body io.ReadCloser, statusCode int, err error) {
	var response *http.Response

	if body != nil {
		response = &http.Response{
			StatusCode: statusCode,
			Body:       body,
		}
	}

	c.DoFunc = func(*http.Request) (*http.Response, error) {
		return response, err
	}
}

// Do is mock client's custom do function
func (c *MockClientImpl) Do(r *http.Request) (*http.Response, error) {
	if c.DoFunc == nil {
		return nil, fmt.Errorf("mock client's Do function is unassigned")
	}
	return c.DoFunc(r)
}
