package utils

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/jackc/pgx/v4"
)

var (
	postgresConnection *pgx.Conn
	connectionContext  = context.Background()
)

func InitializePostgres() error {
	connection, err := pgx.Connect(context.Background(), GetConfig().ConnectionString)
	if err != nil {
		return err
	}

	postgresConnection = connection

	migrationExists, err := CheckMigrationTable()
	if err != nil {
		return err
	}

	if !migrationExists {
		err = CreateMigrationTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckMigrationTable() (bool, error) {
	var doesExists bool
	err := postgresConnection.QueryRow(connectionContext, `SELECT EXISTS (
		SELECT FROM pg_tables
		WHERE tablename  = 'pg_migrate')`).Scan(&doesExists)

	if err != nil {
		return false, err
	}

	return doesExists, nil
}

func CreateMigrationTable() error {
	_, err := postgresConnection.Exec(connectionContext, `CREATE TABLE pg_migrate
	(
			lock        BOOLEAN NOT NULL
										CONSTRAINT pg_migrate_pk
											PRIMARY KEY,
			executed_at TIMESTAMP DEFAULT NOW()
	);
	
	CREATE UNIQUE INDEX pg_migrate_name_uindex
			ON pg_migrate (lock);
			
			INSERT INTO pg_migrate (lock, executed_at) VALUES (false, NULL)`)

	if err != nil {
		return err
	}

	return nil
}

func GetLastMigration() (bool, error) {
	var migrationLock bool
	err := postgresConnection.QueryRow(connectionContext, "SELECT lock FROM pg_migrate").Scan(&migrationLock)

	if err != nil {
		return false, err
	}

	return migrationLock, nil
}

func UpdateMigrationTable(lockState bool) error {
	_, err := postgresConnection.Exec(connectionContext, "UPDATE pg_migrate SET lock = $1, executed_at = NOW()", lockState)

	if err != nil {
		return err
	}

	return err
}

func ExecuteSqlFile(path string) error {
	fileBytes, ioErr := ioutil.ReadFile(filepath.Join(GetConfig().ExecutableDir, path))

	if ioErr != nil {
		return ioErr
	}

	sqlString := string(fileBytes)
	_, err := postgresConnection.Exec(connectionContext, sqlString)

	if err != nil {
		return err
	}

	return nil
}
