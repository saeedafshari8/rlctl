package spring

import (
	"fmt"
	"github.com/rocketlaunchercloud/rlctl/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
)

type SonarQube struct {
	SonarHost                string
	SonarLogin               string
	SonarUserToken           string
	SonarVersion             string
	SonarQualityGateFailMode string //https://github.com/gabrie-allaigre/sonar-gitlab-plugin -> error, warn or none
}

var (
	sonarPropertiesFileName = "sonar-project.properties"
	sonarPropertiesPath     = "spring/sonar-project.properties"

	sonarVersion             = "sonar-version"
	sonarLogin               = "sonar-login"
	sonarUserToken           = "sonar-user-token"
	sonarHost                = "sonar-host"
	sonarQualityGateFailMode = "sonar-quality_gate_fail_mode"

	defaultSonarQubeInstance = SonarQube{
		SonarHost:                "",
		SonarLogin:               "",
		SonarUserToken:           "",
		SonarVersion:             "2.8",
		SonarQualityGateFailMode: "error",
	}
)

func AddSonarFlagsToCommand(cmd *cobra.Command) {
	cmd.Flags().StringP(sonarHost, "", "", defaultSonarQubeInstance.SonarHost)
	cmd.Flags().StringP(sonarUserToken, "", defaultSonarQubeInstance.SonarUserToken, "SonarQuebe user token")
	cmd.Flags().StringP(sonarLogin, "", defaultSonarQubeInstance.SonarLogin, "SonarQuebe login")
	cmd.Flags().StringP(sonarVersion, "", defaultSonarQubeInstance.SonarVersion, "SonarQuebe library version")
	cmd.Flags().StringP(sonarQualityGateFailMode, "", defaultSonarQubeInstance.SonarQualityGateFailMode, "quality_gate_fail_mode (https://github.com/gabrie-allaigre/sonar-gitlab-plugin)")
}

func CreateSonarInstanceFromCommandFlags(cmd *cobra.Command) SonarQube {
	var sonar = SonarQube{}
	sonar.SonarHost = util.GetValue(cmd, sonarHost)
	sonar.SonarLogin = util.GetValue(cmd, sonarLogin)
	sonar.SonarUserToken = util.GetValue(cmd, sonarUserToken)
	sonar.SonarVersion = util.GetValue(cmd, sonarVersion)
	sonar.SonarQualityGateFailMode = util.GetValue(cmd, sonarQualityGateFailMode)
	return sonar
}

func ParseAndSaveSonarQubeFile(projectRoot string, templateData *SpringProjectConfig) (string, error) {
	sonarTemplate, err := util.GetSpringTemplate(sonarPropertiesPath)
	if err != nil {
		return fmt.Sprintf("Unable to parse %s", sonarPropertiesFileName), err
	}
	sonarFilePath := path.Join(projectRoot, sonarPropertiesFileName)
	parsedTemplate, err := util.ParseTemplate(templateData, sonarPropertiesFileName, sonarTemplate)
	if err != nil {
		return fmt.Sprintf("Unable to parse %s", sonarPropertiesFileName), err
	}
	err = ioutil.WriteFile(sonarFilePath, []byte(parsedTemplate), os.ModePerm)
	if err != nil {
		return fmt.Sprintf("Unable to copy %s", sonarPropertiesFileName), err
	}
	return fmt.Sprintf("%s Successfully generated!", sonarPropertiesFileName), nil
}
