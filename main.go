package main

import (
	"ecommerce-platform/cmd"
	logger "ecommerce-platform/logger/log_service_impl"
	"fmt"
)

func main() {
	logger := logger.New()
	logger.Info("Starting Ecommerce Platform Api")
	fmt.Print("hello")
	cmd.Execute()

}
