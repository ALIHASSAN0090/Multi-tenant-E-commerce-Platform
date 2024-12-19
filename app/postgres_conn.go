package app

import (
	"database/sql"
	config "ecommerce-platform/configs"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgres() (*sql.DB, error) {
	dbconnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.AppConfig.DB_USERNAME, config.AppConfig.DB_PASSWORD, config.AppConfig.DB_HOST, config.AppConfig.DB_PORT, config.AppConfig.DB_DATABASE, config.AppConfig.DB_SSL_MODE)

	db, err := sql.Open("postgres", dbconnection)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	return db, nil
}
