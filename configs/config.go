// config/config.go
package config

import (
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"fmt"
	"log"
	"reflect"
)

var AppConfig = models.IConfig{}

func InitializeConfig(c *models.IConfig) error {
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
	if err := InitializeConfig(&AppConfig); err != nil {
		log.Fatalf("Configuration initialization failed: %v", err)
	}
}
