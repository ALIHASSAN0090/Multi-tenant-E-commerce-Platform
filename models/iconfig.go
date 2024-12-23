package models

type IConfig struct {
	RawVars     map[string]string
	APP_ADDRESS string `env:"APP_ADDRESS" required:"true" default:"8005"`
	DB_HOST     string `env:"PG_HOST" required:"true" default:"localhost"`
	DB_PORT     string `env:"PG_PORT" required:"true" default:"5432"`
	DB_DATABASE string `env:"PG_DB" required:"true" default:"gocommerce"`
	DB_USERNAME string `env:"PG_USER" required:"true" default:"postgres"`
	DB_PASSWORD string `env:"PG_PASSWORD" required:"true" default:"password123"`
	DB_SSL_MODE string `env:"DB_SSL_MODE" required:"true" default:"disable"`
	JWT_SECRET  string `env:"JWT_SECRET" required:"true" default:"dflakf@!£@"`
}
