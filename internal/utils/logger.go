package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const LOG_PATH = "./logs"
var logFile *os.File

func CreateLogFile() (*os.File, error) {
	ts := time.Now().Format("20060102_150405")
	logFilePath	:= filepath.Join(LOG_PATH, fmt.Sprintf("setup_%s.txt", ts))

	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)

	return logFile, nil
}

func LogAndDisplay(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	fmt.Println(msg)
	if logFile != nil {
		log.Println(msg)
	}
}