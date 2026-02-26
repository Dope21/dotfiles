package symlink

import (
	"os"
	"os/exec"
)

var RunningOS string = "windows"
var ConflictHandering string = "error" // replace, skip, error

func CreateSymlink(link, target string) error {
	// check the source if it exist. we don't have to check for target since it will auto create
	// check os and create symlink
		// check if it folder of file
		// if the link already exist (check for conflict)
		// if previlege not allow (check for conflict)
	// check error by any mean
}