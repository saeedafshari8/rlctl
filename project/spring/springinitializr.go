package spring

import (
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const (
	Maven                        = "maven-project"
	Gradle                       = "gradle-project"
	Java                         = "java"
	Kotlin                       = "kotlin"
	SpringBootLatestVersion      = "2.2.4.RELEASE"
	springInitializerUrlTemplate = "spring.initializr.tmpl"
)

func GenerateSpringProject(config *SpringProjectConfig) (string, error) {
	springTemplate, err := util.GetSpringTemplate(springInitializerUrlTemplate)
	if err != nil {
		return "", err
	}
	url, err := util.ParseTemplate(config, springInitializerUrlTemplate, springTemplate)
	if err != nil {
		return "", err
	}

	_, err = downloadAndUnzip(&url)
	if err != nil {
		return "", err
	}

	return path.Join(util.OutputDirectory, (*config).Name), nil
}

func downloadAndUnzip(url *string) ([]string, error) {
	request, err := http.NewRequest("GET", *url, nil)
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
