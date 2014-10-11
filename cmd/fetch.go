package cmd

import (
	"fmt"
	"net/url"

	"github.com/codegangsta/cli"

	"github.com/bts-utils/music-tracker/modules/httplib"
)

var CmdFetch = cli.Command{
	Name:  "fetch",
	Usage: "fetch or update each day data",
	Description: `Fetch or update each day data.
`,
	Action: runFetch,
	Flags: []cli.Flag{
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

	date := c.Args().First()
	params := "active=37X8DHpfiimB7PU5y35rfBcg5Vxj2R6umL&archived=&action=export&format=csv&start=" + url.QueryEscape(date) + "&end=" + url.QueryEscape(date)

	body, err := httplib.ResponseBytes("POST", "https://blockchain.info/export-history", params, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
