package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdSummary = cli.Command{
	Name:        "summary",
	Usage:       "show bitshares music pre-sale summary",
	Description: ``,
	Action:      runSummary,
	Flags:       []cli.Flag{},
}

func runSummary(c *cli.Context) {
}
