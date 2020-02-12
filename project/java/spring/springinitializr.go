package spring

import (
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	Maven                        = "maven-project"
	Gradle                       = "gradle-project"
	Java                         = "java"
	SpringBootLatestVersion      = "2.2.1.RELEASE"
	springInitializerUrlTemplate = "spring.initializr.tmpl"
)

type ProjectConfig struct {
	Type                       string
	Language                   string
	SpringBootVersion          string
	Name                       string
	Description                string
	Group                      string
	AppVersion                 string
	AppProtocol                string
	AppHost                    string
	AppPort                    string
	JavaVersion                string
	Database                   string
	EnableJPA                  bool
	EnableLiquibase            bool
	EnableSecurity             bool
	EnableOAuth2               bool
	EnableAzureActiveDirectory bool
	EnableGitLab               bool
	DockerConfig               Docker
	GitLabCIConfig             GitLabCI
}

func DownloadSpringApplication(config ProjectConfig) (string, error) {
	springTemplate, err := util.GetSpringTemplate(springInitializerUrlTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	url, err := util.ParseTemplate(&config, springInitializerUrlTemplate, springTemplate)
	if err != nil {
		return "", err
	}
	files, err := downloadAndUnzip(url)
	if err != nil {
		return "", err
	}
	for _, fileName := range files {
		log.Println(fileName)
	}
	projectPath := path.Join(util.OutputDirectory, config.Name)
	log.Printf("Spring Boot project created successfully under :%s \n", projectPath)
	return projectPath, nil
}

func downloadAndUnzip(url string) ([]string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	ch := make(chan util.ChannelResponse)
	defer close(ch)
	go util.MakeHttpRequest(request, ch)
	channelResponse := <-ch
	if channelResponse.Success {
		fileName, err := util.GenerateTemporaryFileName()
		if err != nil {
			return nil, err
		}
		if err = ioutil.WriteFile(fileName, channelResponse.Data, os.ModePerm); err != nil {
			return nil, err
		}
		files, err := util.Unzip(fileName, util.OutputDirectory)
		// Remove the temp file
		err = os.Remove(fileName)
		if err == nil {
			return nil, err
		}
		return files, err
	}
	return nil, channelResponse.Error
}
