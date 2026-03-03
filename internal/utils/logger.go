package utils

import "fmt"

func LogInfo(message string, isDivide bool) {
	fmt.Println(message)
	if isDivide {
		fmt.Println()
	}
}