package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CreateBackupFile(filePath, backupPath string) error {
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destPath := filepath.Join(backupPath, filepath.Base(sourceFile.Name()))

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