package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Dope21/dotfiles.git/internal/dotfiles"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use: "setup [config]",
	Short: "setup dotfiles from config",
	Long: "symlink and running script to setup dotfiles as define in config",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := args[0]

		configAbsPath, err := filepath.Abs(configPath)
		if err != nil {
			return err
		}

		if _, err := os.Stat(configAbsPath); err != nil {
			return fmt.Errorf("config file %s does not exist", configAbsPath)
		}

		return dotfiles.Setup(configAbsPath)
	},
}