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
	return parseSpringTemplate(gradleTemplateData, gradleBuildTemplate)
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
	return parseSpringTemplate(dockerTemplateData, dockerfileTemplate)
}

func parseSpringTemplate(templateData *ProjectConfig, templateFile string) string {
	dir, err := os.Getwd()
	util.LogAndExit(err, util.EnvironmentError)
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, templateFile))
	util.LogAndExit(err, util.FileNotFound)
	t, err := template.New(templateFile).Parse(string(file))
	util.LogAndExit(err, util.InvalidTemplate)
	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, templateFile, *templateData)
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
