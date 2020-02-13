package spring

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/project/spring"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
	"log"
)

const (
	NameArg  = "name"
	GroupArg = "group"
)

var (
	springProjectConfig spring.SpringProjectConfig

	gitRepositoryUrl string

	SpringCommand = &cobra.Command{
		Use:   "spring",
		Short: "spring command generates a new spring project",
		Long:  `spring command generates a new spring project.`,
		Run: func(cmd *cobra.Command, args []string) {
			initSpringCmdConfig(cmd)

			projectRootPath, err := spring.GenerateSpringProject(&springProjectConfig)
			util.LogAndExit(err, util.NetworkError)

			log.Printf("Spring Boot project created successfully under :%s \n", projectRootPath)

			switch springProjectConfig.Type {
			case spring.Gradle:
				if springProjectConfig.Language == spring.Java {
					spring.OverwriteJavaGradleBuild(&projectRootPath, &springProjectConfig)
				} else if springProjectConfig.Language == spring.Kotlin {

				}
				spring.CreateGradleDockerfile(&projectRootPath, &springProjectConfig)
			}

			spring.ParseAndSaveAppConfigTemplates(projectRootPath, &springProjectConfig)
			spring.ParseAndSaveCiCdFile(projectRootPath, &springProjectConfig)
			spring.SaveK8sTemplates(&projectRootPath, &springProjectConfig)

			util.GitInitNewRepo(projectRootPath)
			util.GitAddAll(projectRootPath)
			util.GitAddRemote(projectRootPath, gitRepositoryUrl)
			util.GitCommit(projectRootPath, "Initial Commit!")
		},
	}
)

func init() {
	initGradleCmdFlags()
}

func initGradleCmdFlags() {
	SpringCommand.Flags().BoolP("azure-ad", "", false, "Enable Azure Active Directory (default false)")
	SpringCommand.Flags().StringP("app-version", "v", "", "Spring boot application version (default is empty and there will not be any version defined for the project)")
	SpringCommand.Flags().StringP("app-port", "", "8080", "Spring boot application port (default is 8080)")
	SpringCommand.Flags().StringP("app-host", "", "localhost", "Spring application base url host (default localhost)")
	SpringCommand.Flags().StringP("app-protocol", "", "http", "Spring application base url protocol (default http")
	SpringCommand.Flags().StringP("container-port", "p", "8080", "Docker exposed port (default is 8080)")
	SpringCommand.Flags().StringP("container-image", "i", "openjdk:11.0.5-jdk-stretch", "Docker exposed port (default is openjdk:11.0.5-jdk-stretch)")
	SpringCommand.Flags().StringP("description", "", "", "Spring application description")
	SpringCommand.Flags().StringP("database", "", "MYSQL", "JPA Database Name (default is MYSQL)")
	SpringCommand.Flags().StringP("docker-registry", "", "dcr.flix.tech/charter/cust", "Docker Registry URL (default is https://index.docker.io/v1)")
	SpringCommand.Flags().StringP("group", "g", "", "Spring application groupId (default is empty)")
	SpringCommand.Flags().BoolP("gitlabci", "", true, "Create .gitlab-ci config (default is true)")
	SpringCommand.Flags().StringArrayP("gitlabci-tags", "", []string{"docker", "autoscaling"}, ".gitlab-ci tags (default is docker,autoscaling)")
	SpringCommand.Flags().StringArrayP("gitlabci-except", "", []string{"schedules"}, ".gitlab-ci except (default is schedules)")
	SpringCommand.Flags().StringP("git-remote", "", "", "git remote repository url")
	SpringCommand.Flags().StringP("java-version", "j", "11", "Gradle (java)sourceCompatibility version (default is 11)")
	SpringCommand.Flags().BoolP("jpa", "", true, "Enable JPA-Hibernate (default is true)")
	SpringCommand.Flags().BoolP("liquibase", "", false, "Enable Liquibase migration (default is false)")
	SpringCommand.Flags().StringP("language", "l", spring.Java, "Spring project language [java | kotlin | groovy] (default is java)")
	SpringCommand.Flags().StringP("name", "", "", "Spring application name")
	SpringCommand.Flags().BoolP("oauth2", "", false, "Enable OAuth2 (default false)")
	SpringCommand.Flags().BoolP("security", "", false, "Enable Spring security (default false)")
	SpringCommand.Flags().StringP("spring-boot-version", "", spring.SpringBootLatestVersion,
		fmt.Sprintf("Spring boot version (default is %s)", spring.SpringBootLatestVersion))
	SpringCommand.Flags().StringP("type", "t", spring.Gradle, "Spring project type [gradle-project | maven-project] (default is gradle-project)")
}

func initSpringCmdConfig(cmd *cobra.Command) {
	//Mandatory flags
	springProjectConfig.Name = util.GetValue(cmd, NameArg)
	util.ValidateRequired(springProjectConfig.Name, NameArg)
	springProjectConfig.Group = util.GetValue(cmd, GroupArg)
	util.ValidateRequired(springProjectConfig.Group, GroupArg)
	gitRepositoryUrl = util.GetValue(cmd, "git-remote")

	//Optional flags
	springProjectConfig.Type = util.GetValue(cmd, "type")
	springProjectConfig.Description = util.GetValue(cmd, "description")
	springProjectConfig.Language = util.GetValue(cmd, "language")
	springProjectConfig.SpringBootVersion = util.GetValue(cmd, "spring-boot-version")
	springProjectConfig.AppVersion = util.GetValue(cmd, "app-version")
	springProjectConfig.JavaVersion = util.GetValue(cmd, "java-version")
	springProjectConfig.DockerConfig.ExposedPort = util.GetValue(cmd, "container-port")
	springProjectConfig.DockerConfig.Image = util.GetValue(cmd, "container-image")
	springProjectConfig.AppProtocol = util.GetValue(cmd, "app-protocol")
	springProjectConfig.AppHost = util.GetValue(cmd, "app-host")
	springProjectConfig.AppPort = util.GetValue(cmd, "app-port")
	springProjectConfig.EnableJPA = util.GetValueBool(cmd, "jpa")
	springProjectConfig.Database = util.GetValue(cmd, "database")
	springProjectConfig.EnableLiquibase = util.GetValueBool(cmd, "liquibase")
	springProjectConfig.EnableSecurity = util.GetValueBool(cmd, "security")
	springProjectConfig.EnableOAuth2 = util.GetValueBool(cmd, "oauth2")
	springProjectConfig.EnableAzureActiveDirectory = util.GetValueBool(cmd, "azure-ad")
	springProjectConfig.EnableGitLab = util.GetValueBool(cmd, "gitlabci")
	springProjectConfig.DockerConfig.RegistryUrl = util.GetValue(cmd, "docker-registry")
	springProjectConfig.GitLabCIConfig.Tags = util.GetValues(cmd, "gitlabci-tags")
	springProjectConfig.GitLabCIConfig.Excepts = util.GetValues(cmd, "gitlabci-except")
}
