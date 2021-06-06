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

func TestClient_Returns_Raw_Branches(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123456")
	defer os.Unsetenv("GITHUB_TOKEN")

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {

		got := mustRead(req.Body)
		want := `{"query":"{repository(owner: \"joesustaric\", name: \"cdi-test-repo\"){id,refs(refPrefix: \"refs/heads/\"){totalCount}}}"}` + "\n"
		assert.Equal(t, got, want)

		w.Header().Set("Content-Type", "application/json")
		resp := `{"data": {"repository": {"id": "someid","refs": {"totalCount": 23}}}}`
		mustWrite(w, resp)
	})
	testClient := &http.Client{Transport: localRoundTripper{handler: mux}}

	githubClient, _ := clients.NewGitHubClient(testClient)

	repo := "https://github.com/joesustaric/cdi-test-repo.git"

	branches, _ := githubClient.RawBranchCount(repo)

	assert.Equal(t, 23, branches)
}

type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
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
