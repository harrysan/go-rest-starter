package main

import (
	"os"

	"finance-tracker/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Finance Tracker"
	app.Version = "v1.0"
	app.Usage = "A finance tracker API service based on golang."
	app.Commands = []*cli.Command{
		cmd.Start(),
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
