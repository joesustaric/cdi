package clients

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	gitHubClient *githubv4.Client
}

func NewGitHubClient() (*GitHubClient, error) {
	tokenVal, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Printf("%s not set\n", "GITHUB_TOKEN")
		return nil, errors.New("GITHUB_TOKEN not set")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tokenVal},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	return &GitHubClient{gitHubClient: client}, nil
}

