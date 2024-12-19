// config/config.go
package config

import (
	"ecommerce-platform/utils"
	"fmt"
	"log"
	"reflect"
)

var AppConfig = IConfig{}

type IConfig struct {
	RawVars     map[string]string
	APP_ADDRESS string `env:"APP_ADDRESS" required:"true" default:"8005"`
	DB_HOST     string `env:"PG_HOST" required:"true" default:"localhost"`
	DB_PORT     string `env:"PG_PORT" required:"true" default:"5432"`
	DB_DATABASE string `env:"PG_DB" required:"true" default:"projectgolang"`
	DB_USERNAME string `env:"PG_USER" required:"true" default:"postgres"`
	DB_PASSWORD string `env:"PG_PASSWORD" required:"true" default:"newpassword"`
	DB_SSL_MODE string `env:"DB_SSL_MODE" required:"true" default:"disable"`
	JWT_SECRET  string `env:"JWT_SECRET" required:"true" default:"dflakf@!£@"`
}

func (c *IConfig) InitializeConfig() error {
	c.RawVars = make(map[string]string)
	v := reflect.ValueOf(c).Elem()
	t := reflect.TypeOf(c).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envTag := field.Tag.Get("env")
		requiredTag := field.Tag.Get("required")
		defaultTag := field.Tag.Get("default")

		if envTag == "" {
			continue
		}

		value := utils.GetEnvValue(envTag, defaultTag)
		if requiredTag == "true" && value == "" {
			return fmt.Errorf("missing required environment variable: %v", envTag)
		}

		c.RawVars[envTag] = value
		if err := utils.SetFieldValue(v.Field(i), value, envTag); err != nil {
			return err
		}
	}

	return nil
}

func InitConfig() {
	if err := AppConfig.InitializeConfig(); err != nil {
		log.Fatalf("Configuration initialization failed: %v", err)
	}
}
