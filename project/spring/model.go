package spring

type SpringProjectConfig struct {
	BuildTool                  string
	Language                   string
	SpringBootVersion          string
	Name                       string
	Description                string
	Group                      string
	Version                    string
	BuildPath                  string
	ServerProtocol             string
	ServerHost                 string
	ServerPort                 string
	JavaSourceCompatibility    string
	JpaDatabase                string
	EnableJPA                  bool
	EnableLiquibase            bool
	EnableSecurity             bool
	EnableOAuth2               bool
	EnableAzureActiveDirectory bool
	EnableGitLabCI             bool
	EnableKafka                bool
	EnableSonar                bool
	EnableJacoco               bool
	SonarQubeConfig            SonarQube
	DockerConfig               Docker
	GitLabCIConfig             GitLabCI
}
