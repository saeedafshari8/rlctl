package spring

import (
	"bytes"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"os"
	"text/template"
)

var (
	gradleBuildTemplate = "project/java/spring/build.gradle.tmpl"
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
