package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "root",
	Short: "Short description",
	Long: "Long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from cobra CLI")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}