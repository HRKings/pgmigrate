package main

import (
	"github.com/urfave/cli/v2"
	"os"

	"github.com/HRKings/pgmigrate/commands"
	"github.com/HRKings/pgmigrate/utils"
)

func main() {
	app := &cli.App{
		Name:  "pgmigrate",
		Usage: "A really simple migration tool for PostgreSQL",
		Before: func(context *cli.Context) error {
			utils.SetConnectionString(context.String("connection"))

			err := utils.InitializePostgres()
			if err != nil {
				return err
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "connection",
				Aliases:  []string{"c"},
				Usage:    "the connection string for PostgreSQL",
				EnvVars:  []string{"PGMIGRATE_CONNECTION_STRING"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			&commands.UpCommand,
			&commands.DownCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
