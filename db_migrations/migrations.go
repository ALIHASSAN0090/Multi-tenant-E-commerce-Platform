package db_migrations

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
)

const MigrationPath = "file://./db_migrations/migration_files"

type Migration struct {
}

func NewMigration() Migration {
	return Migration{}
}

func (m *Migration) RunMigrations(db *sql.DB) error {

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create database driver: %w", err)
	}

	source, err := (&file.File{}).Open(MigrationPath)
	if err != nil {
		return fmt.Errorf("could not open migration files: %w", err)
	}

	migration, err := migrate.NewWithInstance(
		"file", source, "postgres", driver)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Up migration failed: %v. Attempting to run down migrations.", err)
		if downErr := migration.Down(); downErr != nil {
			return fmt.Errorf("could not run down migrations after up failure: %w", downErr)
		}
		if upErr := migration.Up(); upErr != nil && upErr != migrate.ErrNoChange {
			return fmt.Errorf("could not run up migrations after down: %w", upErr)
		}
	}

	log.Println("Migrations applied successfully!")

	return nil
}
