package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

/*
	func fileHash(file) {
		fhash := cli.Command{
			Name:
		}

		if err := cmd.Run(context.Background(), os.Args); err != nil {
			log.Fatal(err)
		}
	}
*/
func textHash(text string) {

	txthash := cli.Command{
		Name:        "txtHash",
		Usage:       "Outputs a hash digest for the inputted text",
		Description: "Outputs a hash digest for the input text, Uses SHA3-256 by default unless otherwise is specified in the flag",
	}

	err := txthash.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	} else {
		return
	}
}
