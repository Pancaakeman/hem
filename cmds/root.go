package cmds

import (
	"hem/cmds/hash"

	"github.com/urfave/cli/v3"
)

func Commands() []*cli.Command {
	commands := []*cli.Command{
		hash.TextHash(),
	}
	return commands
}
