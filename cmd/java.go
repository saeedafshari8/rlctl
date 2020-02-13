package cmd

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/project/java/spring"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
)

type JavaProjectConfig struct {
	SpringProjectConfig spring.ProjectConfig
	GitConfig           util.Git
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

			projectRootPath, err := spring.DownloadSpringApplication(javaProjectConfig.SpringProjectConfig)
			util.LogAndExit(err, util.NetworkError)

			switch javaProjectConfig.SpringProjectConfig.Type {
			case spring.Gradle:
				spring.OverwriteGradleBuild(projectRootPath, spring.ParseGradleTemplate(&javaProjectConfig.SpringProjectConfig))
				spring.CreateDockerfile(projectRootPath, spring.ParseDockerTemplate(&javaProjectConfig.SpringProjectConfig))
				spring.ParseAndSaveAppConfigTemplates(projectRootPath, &javaProjectConfig.SpringProjectConfig)
				spring.ParseAndSaveCiCdFile(projectRootPath, &javaProjectConfig.SpringProjectConfig)
				spring.SaveK8sTemplates(&projectRootPath, &javaProjectConfig.SpringProjectConfig)
			}

			util.GitInitNewRepo(projectRootPath)
			util.GitAddAll(projectRootPath)
			util.GitAddRemote(projectRootPath, javaProjectConfig.GitConfig.RepositoryUrl)
			util.GitCommit(projectRootPath, "Initial Commit!")
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
	cmdJava.Flags().StringP("docker-registry", "", "dcr.flix.tech/charter/cust", "Docker Registry URL (default is https://index.docker.io/v1)")
	cmdJava.Flags().StringP("group", "g", "", "Spring application groupId (default is empty)")
	cmdJava.Flags().BoolP("gitlabci", "", true, "Create .gitlab-ci config (default is true)")
	cmdJava.Flags().StringArrayP("gitlabci-tags", "", []string{"docker", "autoscaling"}, ".gitlab-ci tags (default is docker,autoscaling)")
	cmdJava.Flags().StringArrayP("gitlabci-except", "", []string{"schedules"}, ".gitlab-ci except (default is schedules)")
	cmdJava.Flags().StringP("git-remote", "", "", "git remote repository url")
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

func initJavaConfig(cmd *cobra.Command) {
	//Mandatory flags
	javaProjectConfig.SpringProjectConfig.Name = util.GetValue(cmd, NameArg)
	util.ValidateRequired(javaProjectConfig.SpringProjectConfig.Name, NameArg)
	javaProjectConfig.SpringProjectConfig.Group = util.GetValue(cmd, GroupArg)
	util.ValidateRequired(javaProjectConfig.SpringProjectConfig.Group, GroupArg)
	javaProjectConfig.GitConfig.RepositoryUrl = util.GetValue(cmd, "git-remote")

	//Optional flags
	javaProjectConfig.SpringProjectConfig.Type = util.GetValue(cmd, "type")
	javaProjectConfig.SpringProjectConfig.Description = util.GetValue(cmd, "description")
	javaProjectConfig.SpringProjectConfig.Language = util.GetValue(cmd, "language")
	javaProjectConfig.SpringProjectConfig.SpringBootVersion = util.GetValue(cmd, "spring-boot-version")
	javaProjectConfig.SpringProjectConfig.AppVersion = util.GetValue(cmd, "app-version")
	javaProjectConfig.SpringProjectConfig.JavaVersion = util.GetValue(cmd, "java-version")
	javaProjectConfig.SpringProjectConfig.DockerConfig.ExposedPort = util.GetValue(cmd, "container-port")
	javaProjectConfig.SpringProjectConfig.DockerConfig.Image = util.GetValue(cmd, "container-image")
	javaProjectConfig.SpringProjectConfig.AppProtocol = util.GetValue(cmd, "app-protocol")
	javaProjectConfig.SpringProjectConfig.AppHost = util.GetValue(cmd, "app-host")
	javaProjectConfig.SpringProjectConfig.AppPort = util.GetValue(cmd, "app-port")
	javaProjectConfig.SpringProjectConfig.EnableJPA = util.GetValueBool(cmd, "jpa")
	javaProjectConfig.SpringProjectConfig.Database = util.GetValue(cmd, "database")
	javaProjectConfig.SpringProjectConfig.EnableLiquibase = util.GetValueBool(cmd, "liquibase")
	javaProjectConfig.SpringProjectConfig.EnableSecurity = util.GetValueBool(cmd, "security")
	javaProjectConfig.SpringProjectConfig.EnableOAuth2 = util.GetValueBool(cmd, "oauth2")
	javaProjectConfig.SpringProjectConfig.EnableAzureActiveDirectory = util.GetValueBool(cmd, "azure-ad")
	javaProjectConfig.SpringProjectConfig.EnableGitLab = util.GetValueBool(cmd, "gitlabci")
	javaProjectConfig.SpringProjectConfig.DockerConfig.RegistryUrl = util.GetValue(cmd, "docker-registry")
	javaProjectConfig.SpringProjectConfig.GitLabCIConfig.Tags = util.GetValues(cmd, "gitlabci-tags")
	javaProjectConfig.SpringProjectConfig.GitLabCIConfig.Excepts = util.GetValues(cmd, "gitlabci-except")
}
