package util

import "github.com/gobuffalo/packr/v2"

const (
	JavaSpring = "java-spring"
)

var (
	springTemplatesBox = packr.New(JavaSpring, "../templates")
)

func GetSpringTemplate(templateName string) (string, error) {
	return springTemplatesBox.FindString(templateName)
}
