package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	response := []byte(`{"message":"Repository creation failed.","errors":[{"resource":"Repository","code":"missing_field","field":"name"},{"resource":"Repository","code":"custom","field":"name","message":"name is too short (minimum is 1 character)"}],"documentation_url":"https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)

	var errResponse ErrorResponse
	err := json.Unmarshal(response, &errResponse)

	assert.Nil(t, err)
	if assert.NotNil(t, errResponse) {
		assert.Equal(t, "Repository creation failed.", errResponse.Message)
		assert.Equal(t, "Repository", errResponse.Errors[0].Resource)
		assert.Equal(t, "missing_field", errResponse.Errors[0].Code)
		assert.Equal(t, "name", errResponse.Errors[0].Field)
		assert.Equal(t, "Repository", errResponse.Errors[1].Resource)
		assert.Equal(t, "custom", errResponse.Errors[1].Code)
		assert.Equal(t, "name", errResponse.Errors[1].Field)
		assert.Equal(t, "name is too short (minimum is 1 character)", errResponse.Errors[1].Message)
		assert.Equal(t, "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user", errResponse.DocumentationURL)
	}

}
