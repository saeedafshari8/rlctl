package gradle

import (
	"bytes"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"os"
	"text/template"
)

var (
	gradleBuildTemplate = "project/java/gradle/build.gradle.tmpl"
)

type GradleProjectConfig struct {
	Group               string
	Version             string
	SourceCompatibility string
}

func ParseTemplate(gradleTemplateData GradleProjectConfig) string {
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
