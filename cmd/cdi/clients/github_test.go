package clients_test

import (
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/joesustaric/cdi/cmd/cdi/clients"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestNewGitHubClient_Error_When_No_Token(t *testing.T) {
	_, e := clients.NewGitHubClient(nil)

	assert.NotNil(t, e)
}

func TestNewGitHubClient_No_Error_With_Token(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123456")
	defer os.Unsetenv("GITHUB_TOKEN")

	client, e := clients.NewGitHubClient(nil)

	assert.Nil(t, e)
	assert.NotNil(t, client)
}

func TestRawBranchCount_Returns_Count_Of_Branches(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123456")
	defer os.Unsetenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "put-real-token-in-here-when-re-recording-response"},
	)

	tr := &oauth2.Transport{
		Base:   http.DefaultTransport,
		Source: oauth2.ReuseTokenSource(nil, ts),
	}
	// Start our recorder
	vcrRecorder, err := recorder.NewAsMode(path.Join("fixtures", "github", t.Name()), recorder.ModeReplaying, tr)
	require.NoError(t, err)
	defer vcrRecorder.Stop()

	httpClient := &http.Client{
		Transport: vcrRecorder,
	}

	githubClient, _ := clients.NewGitHubClient(httpClient)

	repo := "https://github.com/joesustaric/cdi-test-repo"

	count, err := githubClient.RawBranchCount(repo)

	require.NoError(t, err)
	assert.Equal(t, 23, count)
}

func TestRawBranchCount_Returns_Error_When_Bad_Token(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123456")
	defer os.Unsetenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "put-real-token-in-here-when-re-recording-response"},
	)

	tr := &oauth2.Transport{
		Base:   http.DefaultTransport,
		Source: oauth2.ReuseTokenSource(nil, ts),
	}
	// Start our recorder
	vcrRecorder, err := recorder.NewAsMode(path.Join("fixtures", "github", t.Name()), recorder.ModeReplaying, tr)
	require.NoError(t, err)
	defer vcrRecorder.Stop()

	httpClient := &http.Client{
		Transport: vcrRecorder,
	}

	githubClient, _ := clients.NewGitHubClient(httpClient)

	repo := "https://github.com/joesustaric/cdi-test-repo.git"

	_, e := githubClient.RawBranchCount(repo)

	assert.NotNil(t, e)
}
