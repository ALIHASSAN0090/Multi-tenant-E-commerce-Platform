package app

import (
	"database/sql"
	config "ecommerce-platform/configs"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgres() (*sql.DB, error) {
	dbconnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.Host, config.Cfg.PGPort, config.Cfg.PGUser, config.Cfg.PGPassword, config.Cfg.PGDB, config.Cfg.SSLMode)

	db, err := sql.Open("postgres", dbconnection)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	return db, nil
}
