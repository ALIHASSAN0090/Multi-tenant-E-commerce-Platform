package main

import (
	"ecommerce-platform/cmd"
	logger "ecommerce-platform/logger/log_service_impl"
)

func main() {
	logger := logger.New()
	logger.Info("Starting Ecommerce Platform Api")

	if err := cmd.Execute(); err != nil {
		logger.Fatal(err)
	}

}
