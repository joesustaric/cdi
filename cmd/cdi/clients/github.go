package clients

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client interface {
	RawBranchCount(link string) (int, error)
}

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

func (c *GitHubClient) RawBranchCount(gitRepoLink string) (int, error) {
	org, repo := c.extractOrgAndRepo(gitRepoLink)

	var query struct {
		Repository struct {
			Id   string
			Refs struct {
				TotalCount int
			} `graphql:"refs(refPrefix: \"refs/heads/\")"`
		} `graphql:"repository(owner: $organisationName, name: $repositoryName)"`
	}
	variables := map[string]interface{}{
		"organisationName": githubv4.String(org),
		"repositoryName":   githubv4.String(repo),
	}

	err := c.gitHubClient.Query(context.Background(), &query, variables)
	if err != nil {
		return 0, err
	}

	return query.Repository.Refs.TotalCount, nil
}

func (c *GitHubClient) extractOrgAndRepo(gitRepoLink string) (string, string) {
	split := strings.Split(gitRepoLink, "/")

	return split[3], strings.Split(split[4], ".git")[0]
}
