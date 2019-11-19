package spring

import (
	"bytes"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"os"
	"text/template"
)

const (
	Maven                   = "maven-project"
	Gradle                  = "gradle-project"
	Java                    = "java"
	SpringBootLatestVersion = "2.2.1.RELEASE"
)

var (
	springInitializerUrlTemplate = "project/java/spring/spring.initializr.tmpl"
)

type ProjectConfig struct {
	Type              string
	Language          string
	SpringBootVersion string
	Name              string
	Description       string
	Group             string
	AppVersion        string
	JavaVersion       string
}

func DownloadSpringApplication(config ProjectConfig) {
	url := compileInitializerUrl(config)
	fmt.Println(url)
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
