package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	//"net/url"

	"github.com/codegangsta/cli"

	"github.com/joliv/spark"

	"github.com/bts-utils/music-tracker/modules/httplib"
)

const NOTES = 5e6

var CmdFetch = cli.Command{
	Name:  "fetch",
	Usage: "fetch or update each day data",
	Description: `Fetch or update each day data.
`,
	Action: runFetch,
	Flags:  []cli.Flag{
	//cli.BoolFlag{"all, a", "Fetch all days data", ""},
	},
}

func runFetch(c *cli.Context) {
	/*
		if c.Bool("all") {
			return
		}

		if len(c.Args()) == 0 {
			return
		}

			date := c.Args().First()
			params := "active=37X8DHpfiimB7PU5y35rfBcg5Vxj2R6umL&archived=&action=export&format=csv&start=" + url.QueryEscape(date) + "&end=" + url.QueryEscape(date)

			body, err := httplib.ResponseBytes("POST", "https://blockchain.info/export-history", params, nil)
	*/

	ts := Transactions{}
	err := httplib.ResponseJSON("GET", "http://www1.agsexplorer.com/ticker/music/btc.json", "", nil, &ts)
	if err != nil {
		fmt.Println(err)
	}

	boring_data := []float64{}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Day\tDate\tBTC\tNotes/BTC\t")
	count := 0
	for i, v := range ts.Txs {
		y, m, d := time.Unix(int64(v[0]), 0).Date()
		date := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		count = i + 1
		fmt.Fprintln(w, fmt.Sprintf("%d\t%s\t%f\t%f\t", count, date, v[1], NOTES/v[1]))
		boring_data = append(boring_data, v[1])
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Total BTC\tAVG BTC\tAVG Notes\t")
	fmt.Fprintln(w, fmt.Sprintf("%f\t%f\t%f\t", float64(ts.Total)/1e8, float64(ts.Avg)/1e8, float64(5e6*count)/(float64(ts.Total)/1e8)))
	fmt.Fprintln(w)
	w.Flush()

	sparkline := spark.Line(boring_data)
	fmt.Println("Sparkline:")
	fmt.Println(sparkline)
	fmt.Println("")
}
