package hash

import (
	"context"
	"crypto/sha256"
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
		UsageText:   "txthash --sha256",
		Description: "Outputs a hash digest for the input text, Uses SHA3-256 by default unless otherwise is specified in the flag",
		Aliases:     []string{"th", "txthash", "texthash", "TxtHash"},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "alg",
				Value: "SHA-256",
				Usage: "Hashing Algorithm to use for the hash",
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			if command.String("alg") == "sha256" {
				hasher := sha256.New()
				hasher.Write([]byte(command.Args().Get(0)))
				log.Printf("%x", hasher.Sum(nil))
			} else {
				log.Println("bad :(")
			}
			return nil
		},
	}
	return &txthash
}
