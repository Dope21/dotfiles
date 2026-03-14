package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var toolName string
var scriptName string

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

	maintainCmd.Flags().StringVarP(&toolName, "tool", "t", "", "tool name")
	maintainCmd.Flags().StringVarP(&scriptName, "script", "s", "", "script name")
	maintainCmd.MarkFlagRequired("tool")
	maintainCmd.MarkFlagRequired("script")
	rootCmd.AddCommand(maintainCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}