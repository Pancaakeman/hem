package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	help := cli.Command{
		Name:        "help",
		Usage:       "Help Command for hem",
		Description: "hem or hash 'em all is a command line utility to produce file hash digests",
	}

	err := help.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	} else {
		return
	}
}
