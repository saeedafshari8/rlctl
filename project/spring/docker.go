package spring

import (
	"github.com/rocketlaunchercloud/rlctl/util"
	"github.com/spf13/cobra"
)

type Docker struct {
	ExposedPort string
	Image       string
	Name        string
	RegistryUrl string
	BashImage   string
}

var (
	containerPort     = "container-port"
	containerImage    = "container-image"
	containerRegistry = "container-registry"
	bashImage         = "container-bash-image"

	defaultDockerInstance = Docker{
		ExposedPort: "8080",
		Image:       "openjdk:11.0.5-jdk-stretch",
		RegistryUrl: "",
		BashImage:   "bash:3.2.57",
	}
)

func AddDockerFlagsToCommand(cmd *cobra.Command) {
	cmd.Flags().StringP(containerPort, "", defaultDockerInstance.ExposedPort, "Docker exposed port")
	cmd.Flags().StringP(containerImage, "", defaultDockerInstance.Image, "Docker exposed port")
	cmd.Flags().StringP(containerRegistry, "", defaultDockerInstance.RegistryUrl, "Docker Registry URL")
	cmd.Flags().StringP(bashImage, "", defaultDockerInstance.BashImage, "Docker Registry URL")
}

func CreateDockerInstanceFromCommandFlags(cmd *cobra.Command) Docker {
	var docker = Docker{}
	docker.ExposedPort = util.GetValue(cmd, containerPort)
	docker.Image = util.GetValue(cmd, containerImage)
	docker.RegistryUrl = util.GetValue(cmd, containerRegistry)
	docker.BashImage = util.GetValue(cmd, bashImage)
	return docker
}
