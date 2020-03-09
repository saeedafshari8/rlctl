package spring

import (
	"fmt"
	"github.com/rocketlaunchercloud/rlctl/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

type GitLabCI struct {
	Tags                    []string
	Excepts                 []string
	K8SDeployStagingEnvTags []string
	K8SDeployProdEnvTags    []string
	Deployer                string
	K8SDevNamespace         string
	K8SProdNamespace        string
	K8SDevCluster           string
	K8SProdCluster          string
	SonarQubeScannerImage   string
}

var (
	moPath           = "buildpipeline/mo.sh"
	gitlabCITemplate = "buildpipeline/.gitlab-ci-default.yml"
	gitlabCI         = ".gitlab-ci.yml"

	gitlabCITags                = "gitlab-ci-tags"
	gitlabCIExcept              = "gitlab-ci-except"
	gitlabCIDeployStagingTags   = "gitlab-ci-k8s-deploy-staging-tags"
	gitlabCIDeployProdTags      = "gitlab-ci-k8s-deploy-prod-tags"
	gitlabCIDeployer            = "gitlab-ci-deployer"
	gitlabCIK8SStagingNamespace = "gitlab-ci-k8s-staging-namespace"
	gitlabCIK8SProdNamespace    = "gitlab-ci-k8s-prod-namespace"
	gitlabCIK8SStagingCluster   = "gitlab-ci-k8s-staging-cluster"
	gitlabCIK8SProdCluster      = "gitlab-ci-k8s-prod-cluster"
	gitlabCISonarScannerImage   = "gitlab-ci-sonar-scanner-image"

	defaultGitlabCIInstance = GitLabCI{
		Tags:                    []string{},
		Excepts:                 []string{},
		K8SDeployStagingEnvTags: []string{},
		K8SDeployProdEnvTags:    []string{},
		Deployer:                "",
		K8SDevNamespace:         "",
		K8SProdNamespace:        "",
		K8SDevCluster:           "",
		K8SProdCluster:          "",
		SonarQubeScannerImage:   "",
	}
)

func AddGitlabCIFlagsToCommand(cmd *cobra.Command) {
	cmd.Flags().StringArrayP(gitlabCITags, "", defaultGitlabCIInstance.Tags, ".gitlab-ci tags")
	cmd.Flags().StringArrayP(gitlabCIExcept, "", defaultGitlabCIInstance.Excepts, ".gitlab-ci except")
	cmd.Flags().StringArrayP(gitlabCIDeployStagingTags, "", defaultGitlabCIInstance.K8SDeployStagingEnvTags, ".gitlab-ci except")
	cmd.Flags().StringArrayP(gitlabCIDeployProdTags, "", defaultGitlabCIInstance.K8SDeployProdEnvTags, ".gitlab-ci except")
	cmd.Flags().StringP(gitlabCIDeployer, "", defaultGitlabCIInstance.Deployer, ".gitlab-ci except")
	cmd.Flags().StringP(gitlabCIK8SStagingNamespace, "", defaultGitlabCIInstance.K8SDevNamespace, ".gitlab-ci except")
	cmd.Flags().StringP(gitlabCIK8SProdNamespace, "", defaultGitlabCIInstance.K8SProdNamespace, ".gitlab-ci except")
	cmd.Flags().StringP(gitlabCIK8SStagingCluster, "", defaultGitlabCIInstance.K8SDevCluster, ".gitlab-ci except")
	cmd.Flags().StringP(gitlabCIK8SProdCluster, "", defaultGitlabCIInstance.K8SProdCluster, ".gitlab-ci except")

	cmd.Flags().StringP(gitlabCISonarScannerImage, "", defaultGitlabCIInstance.SonarQubeScannerImage, "sonar-scanner image")
}

func CreateGitlabCIInstanceFromCommandFlags(cmd *cobra.Command) GitLabCI {
	var ci = GitLabCI{}
	ci.Tags = util.GetValues(cmd, gitlabCITags)
	ci.Excepts = util.GetValues(cmd, gitlabCIExcept)
	ci.K8SDeployStagingEnvTags = util.GetValues(cmd, gitlabCIDeployStagingTags)
	ci.K8SDeployProdEnvTags = util.GetValues(cmd, gitlabCIDeployProdTags)
	ci.Deployer = util.GetValue(cmd, gitlabCIDeployer)
	ci.K8SDevNamespace = util.GetValue(cmd, gitlabCIK8SStagingNamespace)
	ci.K8SProdNamespace = util.GetValue(cmd, gitlabCIK8SProdNamespace)
	ci.K8SDevCluster = util.GetValue(cmd, gitlabCIK8SStagingCluster)
	ci.K8SProdCluster = util.GetValue(cmd, gitlabCIK8SProdCluster)
	ci.SonarQubeScannerImage = util.GetValue(cmd, gitlabCISonarScannerImage)
	return ci
}

func ParseAndSaveCiCdFile(projectRoot string, templateData *SpringProjectConfig) {
	configPath := path.Join(projectRoot, "build_pipeline")
	util.CreateDirIfNotExists(&configPath)

	mo, err := util.GetSpringTemplate(moPath)
	moFilePath := path.Join(configPath, "mo.sh")
	err = ioutil.WriteFile(moFilePath, []byte(mo), os.ModePerm)
	if err != nil {
		util.LogMessageAndExit("Unable to copy mo.sh")
	}

	_, err = exec.Command("chmod", "777", moFilePath).Output()
	if err != nil {
		util.LogMessageAndExit("Unable to make mo.sh executable!")
	}

	templateStr, err := util.GetSpringTemplate(gitlabCITemplate)
	parsedTemplate, err := util.ParseTemplate(templateData, gitlabCI, templateStr)

	filePath := path.Join(projectRoot, gitlabCI)
	err = ioutil.WriteFile(filePath, []byte(parsedTemplate), os.ModePerm)
	if err != nil {
		util.LogMessageAndExit(fmt.Sprintf("Unable to save %s", filePath))
	}
	log.Printf("%s config file created successfully!", gitlabCI)
}
