package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Host               string `mapstructure:"HOST"`
	Port               int    `mapstructure:"PORT"`
	PGPort             int    `mapstructure:"PG_PORT"`
	PGDB               string `mapstructure:"PG_DB"`
	PGUser             string `mapstructure:"PG_USER"`
	PGPassword         string `mapstructure:"PG_PASSWORD"`
	SSLMode            string `mapstructure:"PG_SSL_MODE"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	AdminEmail         string `mapstructure:"ADMIN_DEFAULT_EMAIL"`
	AdminPassword      string `mapstructure:"ADMIN_DEFAULT_PASSWORD"`
	AdminRole          string `mapstructure:"ADMIN_ROLE"`
	GoogleClientID     string `mapstructure:"OAUTH_GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"OAUTH_GOOGLE_CLIENT_SECRET"`
}

var Cfg Config

func LoadConfig() error {
	viper.AddConfigPath("./")
	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName(".env.prod")
	} else if os.Getenv("ENV") == "staging" {
		viper.SetConfigName(".env.staging")
	} else {
		viper.SetConfigName(".env.dev")
	}
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	viper.AutomaticEnv()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Cfg)
	return err
}
