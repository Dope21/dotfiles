package cmd

import (
	"github.com/Dope21/dotfiles.git/internal/dotfiles"
	"github.com/spf13/cobra"
)

var maintainCmd = &cobra.Command{
	Use: "maintain",
	Short: "maintain script for each tool",
	Long: "run a maintain script defined in config for a tool that need manual task",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.Maintain()
	},
}