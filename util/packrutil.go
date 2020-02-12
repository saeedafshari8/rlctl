package util

import "github.com/gobuffalo/packr/v2"

const (
	JavaSpring = "java-spring"
)

var (
	springTemplatesBox         = packr.New(JavaSpring, "../project/java/spring/templates")
	springPipelineTemplatesBox = packr.New(JavaSpring, "../project/java/spring/templates/buildpipeline")
)

func GetSpringTemplate(templateName string) (string, error) {
	return springTemplatesBox.FindString(templateName)
}

func GetSpringPipelineTemplate(templateName string) (string, error) {
	return springPipelineTemplatesBox.FindString(templateName)
}
