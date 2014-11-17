package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdShow = cli.Command{
	Name:        "show",
	Usage:       "show daily donations",
	Description: ``,
	Action:      runShow,
	Flags:       []cli.Flag{},
}

func runShow(c *cli.Context) {
}
