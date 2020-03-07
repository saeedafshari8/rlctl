package cmd

import (
	"github.com/saeedafshari8/flixinit/project/spring"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
	"log"
	"net/url"
)

const (
	group                   = "group"
	name                    = "name"
	description             = "description"
	language                = "language"
	version                 = "version"
	javaSourceCompatibility = "java-source-compatibility"
	buildTool               = "build-tool"
	springBootVersion       = "spring-boot-version"
	serverPort              = "server-port"
	serverHost              = "server-host"
	serverProtocol          = "server-protocol"
	jpaEnabled              = "jpa-enabled"
	jpaDatabase             = "jpa-database"
	liquibaseEnabled        = "liquibase-enabled"
	securityEnabled         = "security-enabled"
	securityOauth2          = "security-oauth2"
	kafkaEnabled            = "kafka-enabled"
	azureEnabled            = "azure-enabled"

	containerPort     = "container-port"
	containerImage    = "container-image"
	containerRegistry = "container-registry"

	gitlabCIEnabled = "gitlab-ci-enabled"
	gitlabCITags    = "gitlab-ci-tags"
	gitlabCIExcept  = "gitlab-ci-except"

	gitRepoUrl = "git-repo-url"
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

			switch springProjectConfig.BuildTool {
			case spring.Gradle:
				if springProjectConfig.Language == spring.Java {
					spring.OverwriteJavaGradleBuild(&projectRootPath, &springProjectConfig)
				} else if springProjectConfig.Language == spring.Kotlin {
					err := spring.OverwriteKotlinGradleBuild(&projectRootPath, &springProjectConfig)
					util.LogAndExit(err, util.NetworkError)
				}
				spring.CreateGradleDockerfile(&projectRootPath, &springProjectConfig)
			}

			spring.ParseAndSaveAppConfigTemplates(projectRootPath, &springProjectConfig)
			spring.ParseAndSaveCiCdFile(projectRootPath, &springProjectConfig)
			spring.SaveK8sTemplates(&projectRootPath, &springProjectConfig)

			if gitRepositoryUrl != "" {
				util.GitInitNewRepo(projectRootPath)
				util.GitAddAll(projectRootPath)
				util.GitAddRemote(projectRootPath, gitRepositoryUrl)
				util.GitCommit(projectRootPath, "Initial Commit!")
				log.Println("Generated files committed to the repository successfully!")
			}
		},
	}
)

func init() {
	initGradleCmdFlags()
}

func initGradleCmdFlags() {
	SpringCommand.Flags().BoolP(azureEnabled, "", false, "Enable Azure Active Directory")
	SpringCommand.Flags().StringP(version, "v", "", "Spring boot application version")
	SpringCommand.Flags().StringP(description, "", "", "Spring application description")
	SpringCommand.Flags().StringP(serverPort, "", "8080", "Spring boot application port")
	SpringCommand.Flags().StringP(serverHost, "", "localhost", "Spring application base url host")
	SpringCommand.Flags().StringP(serverProtocol, "", "http", "Spring application base url protocol")
	SpringCommand.Flags().StringP(jpaDatabase, "", "MYSQL", "JPA Database Name")
	SpringCommand.Flags().StringP(group, "g", "", "Spring application groupId")
	SpringCommand.Flags().StringP(javaSourceCompatibility, "j", "11", "Java source compatibility version")
	SpringCommand.Flags().BoolP(jpaEnabled, "", true, "Enable JPA-Hibernate")
	SpringCommand.Flags().BoolP(liquibaseEnabled, "", false, "Enable Liquibase migration")
	SpringCommand.Flags().StringP(language, "l", spring.Java, "Spring project language [java | kotlin | groovy]")
	SpringCommand.Flags().StringP(name, "", "", "Spring application name")
	SpringCommand.Flags().BoolP(securityOauth2, "", false, "Enable OAuth2")
	SpringCommand.Flags().BoolP(securityEnabled, "", false, "Enable Spring security")
	SpringCommand.Flags().BoolP(kafkaEnabled, "", false, "Enable Kafka integration")
	SpringCommand.Flags().StringP(springBootVersion, "", spring.SpringBootLatestVersion, "Spring boot version")
	SpringCommand.Flags().StringP(buildTool, "", spring.Gradle, "Spring project type [gradle-project | maven-project]")

	SpringCommand.Flags().StringP(containerPort, "", "8080", "Docker exposed port")
	SpringCommand.Flags().StringP(containerImage, "", "openjdk:11.0.5-jdk-stretch", "Docker exposed port")
	SpringCommand.Flags().StringP(containerRegistry, "", "dcr.flix.tech/charter/cust", "Docker Registry URL")

	SpringCommand.Flags().BoolP(gitlabCIEnabled, "", true, "Create .gitlab-ci config")
	SpringCommand.Flags().StringArrayP(gitlabCITags, "", []string{"docker", "autoscaling"}, ".gitlab-ci tags")
	SpringCommand.Flags().StringArrayP(gitlabCIExcept, "", []string{"schedules"}, ".gitlab-ci except")

	SpringCommand.Flags().StringP(gitRepoUrl, "", "", "git remote repository url")
}

func initSpringCmdConfig(cmd *cobra.Command) {
	//Mandatory flags
	springProjectConfig.Name = util.GetValue(cmd, name)
	util.ValidateRequired(springProjectConfig.Name, name)
	springProjectConfig.Group = util.GetValue(cmd, group)
	util.ValidateRequired(springProjectConfig.Group, group)

	//Optional flags
	springProjectConfig.BuildTool = util.GetValue(cmd, buildTool)
	springProjectConfig.Description = url.QueryEscape(util.GetValue(cmd, description))
	springProjectConfig.Language = util.GetValue(cmd, language)
	springProjectConfig.SpringBootVersion = util.GetValue(cmd, springBootVersion)
	springProjectConfig.Version = util.GetValue(cmd, version)
	springProjectConfig.JavaSourceCompatibility = util.GetValue(cmd, javaSourceCompatibility)
	springProjectConfig.ServerProtocol = util.GetValue(cmd, serverProtocol)
	springProjectConfig.ServerHost = util.GetValue(cmd, serverHost)
	springProjectConfig.ServerPort = util.GetValue(cmd, serverPort)
	springProjectConfig.EnableJPA = util.GetValueBool(cmd, jpaEnabled)
	springProjectConfig.JpaDatabase = util.GetValue(cmd, jpaDatabase)
	springProjectConfig.EnableLiquibase = util.GetValueBool(cmd, liquibaseEnabled)
	springProjectConfig.EnableSecurity = util.GetValueBool(cmd, securityEnabled)
	springProjectConfig.EnableOAuth2 = util.GetValueBool(cmd, securityOauth2)
	springProjectConfig.EnableAzureActiveDirectory = util.GetValueBool(cmd, azureEnabled)
	springProjectConfig.EnableKafka = util.GetValueBool(cmd, kafkaEnabled)

	springProjectConfig.DockerConfig.ExposedPort = util.GetValue(cmd, containerPort)
	springProjectConfig.DockerConfig.Image = util.GetValue(cmd, containerImage)
	springProjectConfig.DockerConfig.RegistryUrl = util.GetValue(cmd, containerRegistry)

	springProjectConfig.EnableGitLabCI = util.GetValueBool(cmd, gitlabCIEnabled)
	springProjectConfig.GitLabCIConfig.Tags = util.GetValues(cmd, gitlabCITags)
	springProjectConfig.GitLabCIConfig.Excepts = util.GetValues(cmd, gitlabCIExcept)

	gitRepositoryUrl = util.GetValue(cmd, gitRepoUrl)
}
