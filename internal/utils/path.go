package utils

import (
	"os"
	"path/filepath"
)

func CreatePath(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return nil
}