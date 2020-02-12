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

func GenerateTemporaryFileName() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return path.Join(TempDirectory, uuid.String()+".zip"), nil
}
