package commands

import (
	"fmt"

	"github.com/HRKings/pgmigrate/utils"
	"github.com/urfave/cli/v2"
)

var DownCommand = cli.Command{
	Name:        "down",
	Description: "Executes the migration down",
	Action: func(*cli.Context) error {
		lock, err := utils.GetLastMigration()
		if err != nil {
			return err
		}

		if !lock {
			fmt.Println("No executed migration found")
			return nil
		}

		err = utils.ExecuteSqlFile("down.sql")
		if err != nil {
			return err
		}

		err = utils.UpdateMigrationTable(false)
		if err != nil {
			return err
		}

		fmt.Println("Migration down executed successfully!")
		return nil
	},
}
