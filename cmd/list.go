package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdList = cli.Command{
	Name:        "list",
	Usage:       "",
	Description: ``,
	Action:      runList,
	Flags:       []cli.Flag{},
}

func runList(c *cli.Context) {
}
