package spring

import (
	"bytes"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

const (
	Maven                        = "maven-project"
	Gradle                       = "gradle-project"
	Java                         = "java"
	SpringBootLatestVersion      = "2.2.1.RELEASE"
	springInitializerUrlTemplate = "project/java/spring/spring.initializr.tmpl"
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

func DownloadSpringApplication(config ProjectConfig) string {
	url := compileInitializerUrl(config)
	files, err := get(url)
	util.LogAndExit(err, util.NetworkError)
	for _, fileName := range files {
		log.Println(fileName)
	}
	projectPath := path.Join(util.OutputDirectory, config.Name)
	log.Printf("Spring Boot project created successfully under :%s \n", projectPath)
	return projectPath
}

func get(url string) ([]string, error) {
	ch := make(chan util.ChannelResponse)
	go fetch(url, ch)
	channelResponse := <-ch
	if channelResponse.Success {
		files, err := util.Unzip(channelResponse.Message, util.OutputDirectory)
		// Remove the temp file
		err = os.Remove(channelResponse.Message)
		if err == nil {
			log.Printf("Unable to delete %s\n", channelResponse.Message)
		}
		return files, err
	}
	return nil, channelResponse.Error
}

func fetch(url string, ch chan<- util.ChannelResponse) {
	response, err := http.Get(url)
	if err != nil {
		ch <- util.ChannelResponse{Error: err, Success: false, Message: fmt.Sprint(err)}
		return
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ch <- util.ChannelResponse{Error: err, Success: false, Message: fmt.Sprint(err)}
		return
	}
	fileName := util.GenerateTemporaryFileName()
	if err = ioutil.WriteFile(fileName, responseData, os.ModePerm); err != nil {
		ch <- util.ChannelResponse{Error: err, Success: false, Message: fmt.Sprint(err)}
		return
	}
	ch <- util.ChannelResponse{Error: nil, Success: true, Message: fileName}
}

func compileInitializerUrl(config ProjectConfig) string {
	dir, err := os.Getwd()
	util.LogAndExit(err, util.EnvironmentError)
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, springInitializerUrlTemplate))
	util.LogAndExit(err, util.FileNotFound)
	t, err := template.New(springInitializerUrlTemplate).Parse(string(file))
	util.LogAndExit(err, util.InvalidTemplate)
	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, springInitializerUrlTemplate, config)
	util.LogAndExit(err, util.InvalidTemplate)
	return tmpl.String()
}
