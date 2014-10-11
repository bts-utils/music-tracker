package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdSummary = cli.Command{
	Name:        "summary",
	Usage:       "",
	Description: ``,
	Action:      runSummary,
	Flags:       []cli.Flag{},
}

func runSummary(c *cli.Context) {
}
