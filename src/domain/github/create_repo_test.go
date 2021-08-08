package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequest(t *testing.T) {
	sampleRawRequest := []byte(`{"name":"Hello-World","description":"a temporary repo","private":true,"has_issues":false,"has_projects":false,"has_wiki":true}`)

	newRequest := &CreateRepoRequest{
		Name:        "Hello-World",
		Description: "a temporary repo",
		Private:     true,
		HasWiki:     true,
	}

	js, err := json.Marshal(newRequest)

	assert.Nil(t, err)
	if assert.NotNil(t, js) {
		assert.EqualValues(t, sampleRawRequest, js)
	}
}
func TestCreateRepoResponse(t *testing.T) {
	response := []byte(`{"id":489088539,"name":"Hello-World","full_name":"HarryChopra/Hello-World","private":false,"owner":{"login":"HarryChopra","avatar_url":"https://avatars.githubusercontent.com/u/22281143?v=4","html_url":"https://github.com/HarryChopra","type":"User"},"html_url":"https://github.com/HarryChopra/Hello-World","description":"a temporary repo"}`)

	var repoResponse CreateRepoResponse
	err := json.Unmarshal(response, &repoResponse)

	assert.Nil(t, err)
	if assert.NotNil(t, repoResponse) {
		assert.Equal(t, 489088539, repoResponse.Id)
		assert.Equal(t, "Hello-World", repoResponse.Name)
		assert.Equal(t, "HarryChopra/Hello-World", repoResponse.FullName)
		assert.Equal(t, false, repoResponse.Private)
		assert.Equal(t, "HarryChopra", repoResponse.Owner.Login)
		assert.Equal(t, "https://avatars.githubusercontent.com/u/22281143?v=4", repoResponse.Owner.AvatarUrl)
		assert.Equal(t, "https://github.com/HarryChopra", repoResponse.Owner.HtmlUrl)
		assert.Equal(t, "User", repoResponse.Owner.Type)
		assert.Equal(t, "https://github.com/HarryChopra/Hello-World", repoResponse.HtmlUrl)
		assert.Equal(t, "a temporary repo", repoResponse.Description)
	}
}
