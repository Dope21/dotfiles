package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Dope21/dotfiles.git/internal/types"
	"gopkg.in/yaml.v3"
)

func InitialConfig(configPath string) (types.Config, error) {
	var config types.Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	err = copyConfig(configPath)
	if err != nil {
		return config, err
	}

	return config, nil
}

func copyConfig(path string) error {

	sourceFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	destPath := filepath.Join(configDir, "dotfiles", "config.yaml")
	fmt.Println(destPath)

	if err := CreatePath(destPath); err != nil {
		return err
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return destFile.Sync()
}