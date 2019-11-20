package cmd

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/project/java/spring"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
	"log"
)

type JavaProjectConfig struct {
	SpringProjectConfig spring.ProjectConfig
}

const (
	NameArg  = "name"
	GroupArg = "group"
)

var (
	javaProjectConfig JavaProjectConfig

	cmdJava = &cobra.Command{
		Use:   "java",
		Short: "java command generates a new spring/java project",
		Long:  `java command generates a new spring/java project.`,
		Run: func(cmd *cobra.Command, args []string) {
			initJavaConfig(cmd)

			var template string
			switch javaProjectConfig.SpringProjectConfig.Type {
			case spring.Gradle:
				template = spring.ParseGradleTemplate(javaProjectConfig.SpringProjectConfig)
				log.Println("build.gradle template compiled!")
			}

			projectRootPath := spring.DownloadSpringApplication(javaProjectConfig.SpringProjectConfig)
			spring.OverwriteGradleBuild(projectRootPath, template)
		},
	}
)

func init() {
	cmdJava.Flags().StringP("app-version", "v", "", "Spring boot application version (default is empty and there will not be any version defined for the project)")
	cmdJava.Flags().StringP("description", "", "", "Spring application description")
	cmdJava.Flags().StringP("group", "g", "", "Spring application groupId (default is empty)")
	cmdJava.Flags().StringP("java-version", "j", "11", "Gradle (java)sourceCompatibility version (default is 11)")
	cmdJava.Flags().StringP("language", "l", spring.Java, "Spring project language [java | kotlin | groovy] (default is java)")
	cmdJava.Flags().StringP("name", "", "", "Spring application name")
	cmdJava.Flags().StringP("spring-boot-version", "", spring.SpringBootLatestVersion,
		fmt.Sprintf("Spring boot version (default is %s)", spring.SpringBootLatestVersion))
	cmdJava.Flags().StringP("type", "t", spring.Gradle, "Spring project type [gradle-project | maven-project] (default is gradle-project)")
}

func getValue(cmd *cobra.Command, key string) string {
	s, err := cmd.Flags().GetString(key)
	util.LogAndExit(err, util.ArgMissing)
	return s
}

func initJavaConfig(cmd *cobra.Command) {
	//Mandatory flags
	javaProjectConfig.SpringProjectConfig.Name = getValue(cmd, NameArg)
	checkValue(javaProjectConfig.SpringProjectConfig.Name, NameArg)
	javaProjectConfig.SpringProjectConfig.Group = getValue(cmd, GroupArg)
	checkValue(javaProjectConfig.SpringProjectConfig.Group, GroupArg)

	//Optional flags
	javaProjectConfig.SpringProjectConfig.Type = getValue(cmd, "type")
	javaProjectConfig.SpringProjectConfig.Description = getValue(cmd, "description")
	javaProjectConfig.SpringProjectConfig.Language = getValue(cmd, "language")
	javaProjectConfig.SpringProjectConfig.SpringBootVersion = getValue(cmd, "spring-boot-version")
	javaProjectConfig.SpringProjectConfig.AppVersion = getValue(cmd, "app-version")
	javaProjectConfig.SpringProjectConfig.JavaVersion = getValue(cmd, "java-version")
}

func checkValue(value, key string) {
	if value == "" {
		util.LogMessageAndExit(fmt.Sprintf("%s is mandatory!\n", key))
	}
}
