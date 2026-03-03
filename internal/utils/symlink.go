package utils

import (
	"os"
	"path/filepath"
)

func CreateSymlink(source, link string) error {
	source = os.ExpandEnv(source)
	link = os.ExpandEnv(link)

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