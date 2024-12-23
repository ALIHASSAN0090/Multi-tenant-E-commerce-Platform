package db_migrator

import (
	"database/sql"

	"github.com/pressly/goose"
)

func MigrateDB(db *sql.DB, migrationPath string) error {
	if err := goose.Up(db, migrationPath); err != nil {
		return err
	}

	return nil
}
