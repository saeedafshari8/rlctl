package spring

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	gradleBuildTemplate         = "build.gradle.tmpl"
	kotlinTemplate              = "build.gradle.kts"
	KotlinTemplatePath          = fmt.Sprintf("spring/kotlin/%s", kotlinTemplate)
	dockerfileTemplate          = "Dockerfile.tmpl"
	gradleBuildFileRelativePath = "build.gradle"
	dockerFileRelativePath      = "Dockerfile"
)

func OverwriteJavaGradleBuild(projectRootPath *string, springProjectConfig *SpringProjectConfig) {
	template := parseGradleTemplate(springProjectConfig)

	filePath := path.Join(*projectRootPath, gradleBuildFileRelativePath)
	err := ioutil.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		log.Printf("Unable to overwrite file %s\n", filePath)
	}
	log.Printf("%s updated successfully!", filePath)
}

func OverwriteKotlinGradleBuild(projectRootPath *string, springProjectConfig *SpringProjectConfig) error {
	templateStr, err := util.GetSpringTemplate(KotlinTemplatePath)
	if err != nil {
		return err
	}
	parsedTemplate, err := util.ParseTemplate(springProjectConfig, kotlinTemplate, templateStr)
	if err != nil {
		return err
	}
	filePath := path.Join(*projectRootPath, kotlinTemplate)
	err = ioutil.WriteFile(filePath, []byte(parsedTemplate), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CreateGradleDockerfile(projectRootPath *string, springProjectConfig *SpringProjectConfig) {
	template := parseDockerTemplate(springProjectConfig)

	filePath := path.Join(*projectRootPath, dockerFileRelativePath)
	err := ioutil.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		log.Printf("Unable to write file %s\n", filePath)
	}
	log.Printf("%s updated successfully!", filePath)
}

func parseDockerTemplate(dockerTemplateData *SpringProjectConfig) string {
	springTemplate, err := util.GetSpringTemplate(dockerfileTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	template, err := util.ParseTemplate(dockerTemplateData, dockerfileTemplate, springTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	return template
}

func parseGradleTemplate(gradleTemplateData *SpringProjectConfig) string {
	springTemplate, err := util.GetSpringTemplate(gradleBuildTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	template, err := util.ParseTemplate(gradleTemplateData, gradleBuildTemplate, springTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	return template
}
