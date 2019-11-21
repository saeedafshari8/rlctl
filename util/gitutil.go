package util

import (
	"log"
	"os"
	"os/exec"
)

type Git struct {
	RepositoryUrl string
}

func GitInitNewRepo(repoPath string) {
	cwd, err := os.Getwd()
	LogAndExit(err, EnvironmentError)

	os.Chdir(repoPath)
	cmd := exec.Command("git", "init")
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Print(string(stdout))

	os.Chdir(cwd)
}

func GitAddAll(repoPath string) {
	cwd, err := os.Getwd()
	LogAndExit(err, EnvironmentError)

	os.Chdir(repoPath)
	cmd := exec.Command("git", "add", ".")
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Print(string(stdout))

	os.Chdir(cwd)
}

func GitAddRemote(repoPath string, repositoryUrl string) {
	cwd, err := os.Getwd()
	LogAndExit(err, EnvironmentError)

	os.Chdir(repoPath)
	cmd := exec.Command("git", "remote", "add", "origin", repositoryUrl)
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Print(string(stdout))

	os.Chdir(cwd)
}

func GitCommit(repoPath string, message string) {
	cwd, err := os.Getwd()
	LogAndExit(err, EnvironmentError)

	os.Chdir(repoPath)
	cmd := exec.Command("git", "commit", "-m", message)
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Print(string(stdout))

	os.Chdir(cwd)
}
