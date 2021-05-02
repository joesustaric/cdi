package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md

var version string = "0.1"

func main() {
	app := &cli.App{
		Name:    "cdi",
		Usage:   "cdi <link to repo>",
		Action:  BranchCount,
		Version: version,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func BranchCount(c *cli.Context) error {
	// Branch cound code goes here..
	fmt.Printf("haz %q branches", "123")
	return nil
}
