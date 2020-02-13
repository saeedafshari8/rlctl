package spring

type SpringProjectConfig struct {
	Type                       string
	Language                   string
	SpringBootVersion          string
	Name                       string
	Description                string
	Group                      string
	AppVersion                 string
	AppProtocol                string
	AppHost                    string
	AppPort                    string
	JavaVersion                string
	Database                   string
	EnableJPA                  bool
	EnableLiquibase            bool
	EnableSecurity             bool
	EnableOAuth2               bool
	EnableAzureActiveDirectory bool
	EnableGitLab               bool
	DockerConfig               Docker
	GitLabCIConfig             GitLabCI
}
