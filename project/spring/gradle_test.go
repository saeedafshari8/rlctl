package spring_test

import "testing"
import "github.com/rocketlaunchercloud/rlctl/project/spring"

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
