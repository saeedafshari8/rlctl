package spring

import (
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const (
	gradleBuildTemplate         = "build.gradle.tmpl"
	dockerfileTemplate          = "Dockerfile.tmpl"
	gradleBuildFileRelativePath = "build.gradle"
	dockerFileRelativePath      = "Dockerfile"
)

func ParseGradleTemplate(gradleTemplateData *ProjectConfig) string {
	springTemplate, err := util.GetSpringTemplate(gradleBuildTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	template, err := util.ParseTemplate(gradleTemplateData, gradleBuildTemplate, springTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	return template
}

func ParseDockerTemplate(dockerTemplateData *ProjectConfig) string {
	springTemplate, err := util.GetSpringTemplate(dockerfileTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	template, err := util.ParseTemplate(dockerTemplateData, dockerfileTemplate, springTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	return template
}

func OverwriteGradleBuild(projectRootPath, template string) {
	filePath := path.Join(projectRootPath, gradleBuildFileRelativePath)
	err := ioutil.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		log.Printf("Unable to overwrite file %s\n", filePath)
	}
	log.Printf("%s updated successfully!", filePath)
}

func CreateDockerfile(projectRootPath, template string) {
	filePath := path.Join(projectRootPath, dockerFileRelativePath)
	err := ioutil.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		log.Printf("Unable to write file %s\n", filePath)
	}
	log.Printf("%s updated successfully!", filePath)
}
