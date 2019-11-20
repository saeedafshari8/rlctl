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

			projectRootPath := spring.DownloadSpringApplication(javaProjectConfig.SpringProjectConfig)

			switch javaProjectConfig.SpringProjectConfig.Type {
			case spring.Gradle:
				spring.OverwriteGradleBuild(projectRootPath, spring.ParseGradleTemplate(&javaProjectConfig.SpringProjectConfig))
				spring.CreateDockerfile(projectRootPath, spring.ParseDockerTemplate(&javaProjectConfig.SpringProjectConfig))
				spring.ParseAndSaveAppConfigTemplates(projectRootPath, &javaProjectConfig.SpringProjectConfig)
				log.Println("build.gradle template compiled!")
			}
		},
	}
)

func init() {
	cmdJava.Flags().BoolP("azure-ad", "", false, "Enable Azure Active Directory (default false)")
	cmdJava.Flags().StringP("app-version", "v", "", "Spring boot application version (default is empty and there will not be any version defined for the project)")
	cmdJava.Flags().StringP("app-port", "", "8080", "Spring boot application port (default is 8080)")
	cmdJava.Flags().StringP("app-host", "", "localhost", "Spring application base url host (default localhost)")
	cmdJava.Flags().StringP("app-protocol", "", "http", "Spring application base url protocol (default http")
	cmdJava.Flags().StringP("container-port", "p", "8080", "Docker exposed port (default is 8080)")
	cmdJava.Flags().StringP("container-image", "i", "openjdk:11.0.5-jdk-stretch", "Docker exposed port (default is openjdk:11.0.5-jdk-stretch)")
	cmdJava.Flags().StringP("description", "", "", "Spring application description")
	cmdJava.Flags().StringP("database", "", "MYSQL", "JPA Database Name (default is MYSQL)")
	cmdJava.Flags().StringP("group", "g", "", "Spring application groupId (default is empty)")
	cmdJava.Flags().StringP("java-version", "j", "11", "Gradle (java)sourceCompatibility version (default is 11)")
	cmdJava.Flags().BoolP("jpa", "", true, "Enable JPA-Hibernate (default is true)")
	cmdJava.Flags().BoolP("liquibase", "", false, "Enable Liquibase migration (default is false)")
	cmdJava.Flags().StringP("language", "l", spring.Java, "Spring project language [java | kotlin | groovy] (default is java)")
	cmdJava.Flags().StringP("name", "", "", "Spring application name")
	cmdJava.Flags().BoolP("oauth2", "", false, "Enable OAuth2 (default false)")
	cmdJava.Flags().BoolP("security", "", false, "Enable Spring security (default false)")
	cmdJava.Flags().StringP("spring-boot-version", "", spring.SpringBootLatestVersion,
		fmt.Sprintf("Spring boot version (default is %s)", spring.SpringBootLatestVersion))
	cmdJava.Flags().StringP("type", "t", spring.Gradle, "Spring project type [gradle-project | maven-project] (default is gradle-project)")
}

func getValue(cmd *cobra.Command, key string) string {
	s, err := cmd.Flags().GetString(key)
	util.LogAndExit(err, util.ArgMissing)
	return s
}

func getValueBool(cmd *cobra.Command, key string) bool {
	b, err := cmd.Flags().GetBool(key)
	util.LogAndExit(err, util.ArgMissing)
	return b
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
	javaProjectConfig.SpringProjectConfig.DockerConfig.ExposedPort = getValue(cmd, "container-port")
	javaProjectConfig.SpringProjectConfig.DockerConfig.Image = getValue(cmd, "container-image")
	javaProjectConfig.SpringProjectConfig.AppProtocol = getValue(cmd, "app-protocol")
	javaProjectConfig.SpringProjectConfig.AppHost = getValue(cmd, "app-host")
	javaProjectConfig.SpringProjectConfig.AppPort = getValue(cmd, "app-port")
	javaProjectConfig.SpringProjectConfig.EnableJPA = getValueBool(cmd, "jpa")
	javaProjectConfig.SpringProjectConfig.Database = getValue(cmd, "database")
	javaProjectConfig.SpringProjectConfig.EnableLiquibase = getValueBool(cmd, "liquibase")
	javaProjectConfig.SpringProjectConfig.EnableSecurity = getValueBool(cmd, "security")
	javaProjectConfig.SpringProjectConfig.EnableOAuth2 = getValueBool(cmd, "oauth2")
	javaProjectConfig.SpringProjectConfig.EnableAzureActiveDirectory = getValueBool(cmd, "azure-ad")
}

func checkValue(value, key string) {
	if value == "" {
		util.LogMessageAndExit(fmt.Sprintf("%s is mandatory!\n", key))
	}
}
