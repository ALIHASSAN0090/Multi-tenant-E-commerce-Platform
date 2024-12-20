package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() error {
	return rootCmd.Execute()
}

func Init() {
	fmt.Println(14)
	rootCmd.AddCommand(ApiServerCommand)
}

var rootCmd = &cobra.Command{
	Use: "ecommerce-platform",
}
