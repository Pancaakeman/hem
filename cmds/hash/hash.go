package hash

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"
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

func TextHash() *cli.Command {
	txthash := cli.Command{
		Name:        "txtHash",
		Usage:       "Outputs a hash digest for the inputted text",
		Description: "Outputs a hash digest for the input text, Uses SHA3-256 by default unless otherwise is specified in the flag",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "sha-256",
				Value: "SHA-256",
				Usage: "Hashing Algorithm to use for the hash",
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			if command.String("hash_alg") == "sha-256" {
				log.Println("sha-256")
			} else {
				log.Println("bad :(")
			}
			return nil
		},
	}
	log.Println(txthash)
	return &txthash
}
