package spring_test

import "testing"
import "github.com/saeedafshari8/flixinit/project/spring"

func TestOverwriteKotlinGradleBuild(t *testing.T) {
	var springProjectConfig spring.SpringProjectConfig
	springProjectConfig.Name = "test"
	springProjectConfig.EnableKafka = true

	rootPath := "/tmp"
	err := spring.OverwriteKotlinGradleBuild(&rootPath, &springProjectConfig)
	if err != nil {
		t.Fatal("error happened", springProjectConfig)
	}
}
