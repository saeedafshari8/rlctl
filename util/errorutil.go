package util

import (
	"log"
	"os"
)

const (
	EnvironmentError = 1
	FileNotFound     = 2
	InvalidTemplate  = 3
	ArgMissing       = 4
	NetworkError     = 5
)

type ChannelResponse struct {
	Success bool
	Error   error
	Data    []byte
}

func Error(err error) bool {
	if err != nil {
		log.Fatal(err)
		return true
	}
	return false
}

func LogAndExit(err error, exitCode int) {
	if Error(err) {
		os.Exit(exitCode)
	}
}

func LogMessageAndExit(message string) {
	log.Fatal(message)
}
