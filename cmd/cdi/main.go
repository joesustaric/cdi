package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joesustaric/cdi/cmd/cdi/clients"
	"github.com/urfave/cli/v2"
)

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md

var version = "0.1"

var client clients.Client = nil

func main() {
	app := &cli.App{
		Name:    "cdi",
		Usage:   "cdi <link to repo>",
		Action:  BranchCount,
		Version: version,
	}

	_, ok := os.LookupEnv("TESTING")
	if ok {
		client = clients.NewTestClient()
	} else {
		client, _ = clients.NewGitHubClient(nil)
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func BranchCount(c *cli.Context) error {
	a, b := client.RawBranchCount(c.Args().Get(0))
	fmt.Printf("haz %d branches\n", a)
	return b
}
