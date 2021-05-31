package main

// cd ito this dir and run
// ensure you install the deps from the root
// $ go run main.go

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// I made a test repo here to could branches and ro random things too
// https://github.com/joesustaric/cdi-test-repo
// its public but the Oauth client seems to want a token :shrug:

// This struct below is representing this graph ql query
// {
//   repository(owner: "joesustaric", name: "cdi-test-repo") {
//     id
//     refs(refPrefix: "refs/heads/") {
//       totalCount
//     }
//   }
// }
// Head refs seem to be what we want as "branch count"
// See https://github.com/shurcooL/githubv4#arguments-and-variables

var query struct {
	Repository struct {
		Description string
		Refs        struct {
			TotalCount int
		} `graphql:"refs(refPrefix: \"refs/heads/\")"`
	} `graphql:"repository(owner: \"joesustaric\", name: \"cdi-test-repo\")"`
}

func main() {
	src := oauth2.StaticTokenSource(
		// Put your OAuth token in there or pluck it from the ENV
		&oauth2.Token{AccessToken: "your token from github here"},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Github enterprise
	// client := githubv4.NewEnterpriseClient(os.Getenv("GITHUB_ENDPOINT"), httpClient)

	client := githubv4.NewClient(httpClient)

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}
	fmt.Printf("Total Branches = %d", query.Repository.Refs.TotalCount)
}
