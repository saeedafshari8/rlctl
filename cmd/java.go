package cmd

import (
	"github.com/saeedafshari8/flixinit/project/java/gradle"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
	"log"
)

type JavaProjectConfig struct {
	GradleConfig gradle.GradleProjectConfig
}

var (
	javaProjectConfig JavaProjectConfig

	cmdJava = &cobra.Command{
		Use:   "java",
		Short: "java command generates a new gradle/java project",
		Long:  `java command generates a new gradle/java project.`,
		Run: func(cmd *cobra.Command, args []string) {
			parseGradleTemplate(cmd)
			log.Println("build.gradle template compiled!")
		},
	}
)

func parseGradleTemplate(cmd *cobra.Command) string {
	javaProjectConfig.GradleConfig = buildGradleConfig(cmd)
	return gradle.ParseTemplate(javaProjectConfig.GradleConfig)
}

func init() {
	cmdJava.Flags().StringP("app-version", "v", "", "Gradle application version (default is empty)")
	cmdJava.Flags().StringP("java-version", "j", "11", "Gradle (java)sourceCompatibility version (default is 11)")
	cmdJava.Flags().StringP("group", "g", "", "Gradle project group (default is empty)")
}

func buildGradleConfig(cmd *cobra.Command) gradle.GradleProjectConfig {
	var config gradle.GradleProjectConfig

	s, err := cmd.Flags().GetString("group")
	config.Group = s
	if err != nil || config.Group == "" {
		util.LogMessageAndExit("Gradle project group name is mandatory")
	}

	return config
}
