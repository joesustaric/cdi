package clients_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joesustaric/cdi/cmd/cdi/clients"
	"github.com/stretchr/testify/assert"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {

		got := mustRead(req.Body)
		want := `{"query":"query($organisationName:String!$repositoryName:String!){repository(owner: $organisationName, name: $repositoryName){id,refs(refPrefix: \"refs/heads/\"){totalCount}}}","variables":{"organisationName":"some-org","repositoryName":"some-repo"}}` + "\n"
		assert.Equal(t, want, got)

		w.Header().Set("Content-Type", "application/json")
		resp := `{"data": {"repository": {"id": "someid","refs": {"totalCount": 23}}}}`
		mustWrite(w, resp)
	})
	testClient := &http.Client{Transport: localTestServer{handler: mux}}

	githubClient, _ := clients.NewGitHubClient(testClient)

	repo := "https://github.com/some-org/some-repo.git"

	branches, _ := githubClient.RawBranchCount(repo)

	assert.Equal(t, 23, branches)
}

func TestRawBranchCount_Returns_Error_When_500(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123456")
	defer os.Unsetenv("GITHUB_TOKEN")

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		mustWrite(w, "")
	})
	testClient := &http.Client{Transport: localTestServer{handler: mux}}

	githubClient, _ := clients.NewGitHubClient(testClient)

	repo := "https://github.com/joesustaric/cdi-test-repo.git"

	_, e := githubClient.RawBranchCount(repo)

	assert.NotNil(t, e)
}

type localTestServer struct {
	handler http.Handler
}

func (l localTestServer) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)
	return w.Result(), nil
}

func mustRead(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func mustWrite(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		panic(err)
	}
}
