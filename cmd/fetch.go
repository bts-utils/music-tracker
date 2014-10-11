package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/bts-utils/music-tracker/modules/fetch"
)

var CmdFetch = cli.Command{
	Name:  "fetch",
	Usage: "Fetch or update each day data",
	Description: `Fetch or update each day data.
`,
	Action: runFetch,
	Flags:  []cli.Flag{
		cli.BoolFlag{"all, a", "Fetch all days data", ""},
	},
}

func runFetch(c *cli.Context) {
	if c.Bool("all") {
		return
	}

	if len(c.Args()) == 0 {
		return
	}
}
