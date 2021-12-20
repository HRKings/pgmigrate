package main

import (
	"os"

	"github.com/HRKings/pgmigrate/commands"
	"github.com/HRKings/pgmigrate/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pgmigrate",
		Usage: "A really migration tool for PostgreSQL",
		Commands: []*cli.Command{
			&commands.UpCommand,
			&commands.DownCommand,
		},
	}

	err := utils.InitializePostgres()
	if err != nil {
		panic(err)
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
