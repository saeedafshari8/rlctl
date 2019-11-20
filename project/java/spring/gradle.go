package spring

import (
	"bytes"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

const (
	gradleBuildTemplate         = "project/java/spring/build.gradle.tmpl"
	dockerfileTemplate          = "project/java/spring/Dockerfile.tmpl"
	gradleBuildFileRelativePath = "build.gradle"
	dockerFileRelativePath      = "Dockerfile"
)

func ParseGradleTemplate(gradleTemplateData *ProjectConfig) string {
	dir, err := os.Getwd()

	util.LogAndExit(err, util.EnvironmentError)

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, gradleBuildTemplate))

	util.LogAndExit(err, util.FileNotFound)

	t, err := template.New(gradleBuildTemplate).Parse(string(file))

	util.LogAndExit(err, util.InvalidTemplate)

	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, gradleBuildTemplate, *gradleTemplateData)

	util.LogAndExit(err, util.InvalidTemplate)

	return tmpl.String()
}

func OverwriteGradleBuild(projectRootPath, template string) {
	filePath := path.Join(projectRootPath, gradleBuildFileRelativePath)
	err := ioutil.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		log.Printf("Unable to overwrite file %s\n", filePath)
	}
	log.Printf("%s updated successfully!", filePath)
}

func ParseDockerTemplate(dockerTemplateData *ProjectConfig) string {
	dir, err := os.Getwd()

	util.LogAndExit(err, util.EnvironmentError)

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, dockerfileTemplate))

	util.LogAndExit(err, util.FileNotFound)

	t, err := template.New(dockerfileTemplate).Parse(string(file))

	util.LogAndExit(err, util.InvalidTemplate)

	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, dockerfileTemplate, *dockerTemplateData)

	util.LogAndExit(err, util.InvalidTemplate)

	return tmpl.String()
}

func CreateDockerfile(projectRootPath, template string) {
	filePath := path.Join(projectRootPath, dockerFileRelativePath)
	err := ioutil.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		log.Printf("Unable to write file %s\n", filePath)
	}
	log.Printf("%s updated successfully!", filePath)
}
