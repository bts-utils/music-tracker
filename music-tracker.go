package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/bts-utils/music-tracker/cmd"
)

const APP_VER = "0.0.0.1011"

func init() {
}

func main() {
	app := cli.NewApp()
	app.Name = "music-tracker"
	app.Usage = "BitShares Music Note pre-sale tracker"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdFetch,
		cmd.CmdShow,
		cmd.CmdList,
		cmd.CmdSummary,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
