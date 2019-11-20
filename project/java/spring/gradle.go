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
	gradleBuildFileRelativePath = "build.gradle"
)

func ParseGradleTemplate(gradleTemplateData ProjectConfig) string {
	dir, err := os.Getwd()

	util.LogAndExit(err, util.EnvironmentError)

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, gradleBuildTemplate))

	util.LogAndExit(err, util.FileNotFound)

	t, err := template.New(gradleBuildTemplate).Parse(string(file))

	util.LogAndExit(err, util.InvalidTemplate)

	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, gradleBuildTemplate, gradleTemplateData)

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
