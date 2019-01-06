package main

import (
	"fmt"
	"os"

	"github.com/TukioHQ/togo/actions"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "togo"
	app.Usage = "togo provides tools to convert files to go"
	app.Version = "1.0.0"
	app.Author = "bradrydzewski"
	app.Commands = []cli.Command{
		actions.DDLCommands,
		actions.SQLCommand,
		actions.HTTPCommand,
		actions.HTTPTestCommand,
		actions.TMPLCommand,
		actions.I18nCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
