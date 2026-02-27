package symlink

import "os"

func CreateSymlink(source, link string) error {
	_, err := os.Stat(source)
	if err != nil {
		return err
	}

	linkInfo, err := os.Lstat(link)
	if err != nil {
		return err
	}

	if linkInfo.Mode()&os.ModeSymlink != 0 {
		err = os.Remove(link)
		if err != nil {
			return err
		}
	}

	err = os.Symlink(source, link)
	if err != nil {
		return err
	}

	return nil
}