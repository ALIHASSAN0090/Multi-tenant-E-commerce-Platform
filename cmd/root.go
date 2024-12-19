package cmd

import "github.com/spf13/cobra"

func Execute() error {
	return rootCmd.Execute()
}

func Init() {
	rootCmd.AddCommand(ApiServerCommand)
}

var rootCmd = &cobra.Command{
	Use: "ecommerce-platform",
}
