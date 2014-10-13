package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdCountdown = cli.Command{
	Name:        "countdown",
	Usage:       "countdown to last day",
	Description: ``,
	Action:      runCountdown,
	Flags:       []cli.Flag{},
}

func runCountdown(c *cli.Context) {
}
