package utils

import (
	"fmt"
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