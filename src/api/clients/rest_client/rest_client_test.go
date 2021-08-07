package restclient

import (
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/harrychopra/go-github-microservice/src/api/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	Client = new(mocks.MockClientImpl)
	os.Exit(m.Run())
}

func TestPost_marshalError(t *testing.T) {
	response, err := Post(nil, "", math.NaN)

	assert.Nil(t, response)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "could not marshal")
	}
}

func TestPost_requestCreationError(t *testing.T) {
	body := io.NopCloser(strings.NewReader(`{"id":123}`))

	Client.(*mocks.MockClientImpl).SetDoFunc(body, http.StatusCreated, nil)

	response, err := Post(nil, "http://valid/url", nil)

	assert.Nil(t, err)
	if assert.NotNil(t, response) {
		assert.Equal(t, http.StatusCreated, response.StatusCode)
	}
}
