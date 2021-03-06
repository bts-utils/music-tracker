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

	"github.com/go-libs/ansi/styles"
)

// Each day 500,000 Notes
const DAY_NOTES = 5e6

// 100,000,000
const RATIO = 1e8

var CmdFetch = cli.Command{
	Name:        "fetch",
	Usage:       "fetch or update the latest status",
	Description: `Fetch or update the latest status.`,
	Action:      runFetch,
	Flags:       []cli.Flag{
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

	var (
		boring     []float64
		i          int
		t          Tx
		ty, tm, td = time.Now().UTC().Date()
	)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Day\tDate\tBTC\tNotes/BTC")
	for i, t = range ts.Txs {
		t0, t1 := int64(t[0]), t[1]
		y, m, d := time.Unix(t0, 0).Date()
		date := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		str := fmt.Sprintf("%d\t%s\t%f\t%f", i+1, date, t1, DAY_NOTES/t1)
		if ty == y && tm == m && td == d {
			str = styles.Red.Print(str)
		}
		fmt.Fprintln(w, str)
		boring = append(boring, t1)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "\tTotal BTC\tAVG BTC\tAVG Notes")
	fmt.Fprintln(w, styles.Blue.Print(fmt.Sprintf("\t%f\t%f\t%f", ts.Total/RATIO, ts.Avg/RATIO, float64(DAY_NOTES*(i+1))/(ts.Total/RATIO))))
	fmt.Fprintln(w)

	sparkline := spark.Line(boring)
	fmt.Fprintln(w, "\tSparkline:")
	fmt.Fprintln(w, "\t"+styles.Green.Print(sparkline)+"\n")
	w.Flush()
}
