package main

import (
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

var (
	Version = ""

	startCmd = cli.Command{
		Name:   "start",
		Usage:  "penquin start",
		Action: nil,
		Flags:  nil,
	}
)

func main() {
	penquin := newPenQuin()
	if err := penquin.Run(os.Args); err != nil {
		panic(err)
	}
}

func newPenQuin() *cli.App {
	app := cli.NewApp()
	app.Name = "PenQuin"
	app.Version = Version
	app.Compiled = time.Now()
	app.Copyright = "(c) " + strconv.Itoa(time.Now().Year()) + " Entertainment Venue"
	app.Usage = "PenQuin is a DelayQueue base on REDIS."
	app.Flags = startCmd.Flags
	//commands
	app.Commands = []cli.Command{
		startCmd,
	}
	//action
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			return cli.ShowAppHelp(c)
		}

		return startCmd.Action.(func(c *cli.Context) error)(c)
	}

	return app
}
