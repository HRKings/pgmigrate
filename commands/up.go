package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"

	"github.com/HRKings/pgmigrate/utils"
)

var UpCommand = cli.Command{
	Name:        "up",
	Description: "Executes the migration up",
	Action: func(ctx *cli.Context) error {
		lock, err := utils.GetLastMigration()
		if err != nil {
			return err
		}

		if lock {
			fmt.Println("Migration already executed")
			return nil
		}

		err = utils.ExecuteSqlFile("up.sql")
		if err != nil {
			return err
		}

		err = utils.UpdateMigrationTable(true)
		if err != nil {
			return err
		}

		fmt.Println("Migration up executed successfully!")
		return nil
	},
}
