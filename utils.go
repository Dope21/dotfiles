package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetPlatform() (string, error) {
	os := runtime.GOOS

	switch os {
	case "windows", "linux", "darwin":
		return os, nil
	default:
		return "", fmt.Errorf("%s OS is not supported", runtime.GOOS)
	}
}

func CreateSymlink(source, link string) error {
	sourceAbs, err := filepath.Abs(source)
	if err != nil {
		return err
	}

	linkInfo, err := os.Lstat(link)

	if err == nil && linkInfo.Mode()&os.ModeSymlink != 0 {
		if err = os.Remove(link); err != nil {
			return err
		}
	}

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		if err = CreatePath(link); err != nil {
			return err
		}
	}

	if err = os.Symlink(sourceAbs, link); err != nil {
		return err
	}

	return nil
}

func CreatePath(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return nil
}