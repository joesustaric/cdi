package clients_test

import (
	"os"
	"testing"

	"github.com/joesustaric/cdi/cmd/cdi/clients"
	"github.com/stretchr/testify/assert"
)

func TestNewGitHubClient_Error_When_No_Token(t *testing.T) {

	_, e := clients.NewGitHubClient()

	assert.NotNil(t, e)
}

func TestNewGitHubClient_No_Error_With_Token(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123456")

	client, e := clients.NewGitHubClient()

	assert.Nil(t, e)
	assert.NotNil(t, client)

	os.Unsetenv("GITHUB_TOKEN")
}

