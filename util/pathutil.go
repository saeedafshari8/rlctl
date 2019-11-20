package util

import (
	"github.com/google/uuid"
	"os"
	"path"
)

var (
	TempDirectory   string
	OutputDirectory string
)

func init() {
	dir, err := os.Getwd()
	LogAndExit(err, EnvironmentError)
	TempDirectory = path.Join(dir, "tmp")
	CreateDirIfNotExists(&TempDirectory)
	OutputDirectory = path.Join(dir, "build")
	CreateDirIfNotExists(&OutputDirectory)
}

func CreateDirIfNotExists(dir *string) {
	if result, _ := Exists(*dir); !result {
		LogAndExit(os.Mkdir(*dir, os.ModePerm), EnvironmentError)
	}
}

func GenerateTemporaryFileName() string {
	uuid, e := uuid.NewUUID()
	if e != nil {
		LogMessageAndExit("UUID creation failed!")
	}
	return path.Join(TempDirectory, uuid.String()+".zip")
}
