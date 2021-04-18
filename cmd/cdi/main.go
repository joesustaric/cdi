package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
