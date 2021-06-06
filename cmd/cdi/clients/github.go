package clients

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	gitHubClient *githubv4.Client
}

func NewGitHubClient(httpClient *http.Client) (*GitHubClient, error) {
	tokenVal, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Printf("%s not set\n", "GITHUB_TOKEN")
		return nil, errors.New("GITHUB_TOKEN not set")
	}

	if httpClient == nil {
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: tokenVal},
		)
		httpClient = oauth2.NewClient(context.Background(), src)
	}

	client := githubv4.NewClient(httpClient)

	return &GitHubClient{gitHubClient: client}, nil
}

func (c *GitHubClient) RawBranchCount(repo string) (int, error) {
	var query struct {
		Repository struct {
			Id   string
			Refs struct {
				TotalCount int
			} `graphql:"refs(refPrefix: \"refs/heads/\")"`
		} `graphql:"repository(owner: \"joesustaric\", name: \"cdi-test-repo\")"`
	}

	err := c.gitHubClient.Query(context.Background(), &query, nil)
	if err != nil {
		return 0, err
	}
	return query.Repository.Refs.TotalCount, nil
}
