package spring

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

var (
	moPath           = "buildpipeline/mo.sh"
	gitignorePath    = "buildpipeline/.gitignore.tmpl"
	gitlabCITemplate = "buildpipeline/.gitlab-ci-default.yml"
	gitlabCI         = ".gitlab-ci.yml"
)

type GitLabCI struct {
	Tags    []string
	Excepts []string
}

func ParseAndSaveCiCdFile(projectRoot string, templateData *SpringProjectConfig) {
	if (*templateData).EnableGitLabCI {
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

		gitignore, err := util.GetSpringTemplate(gitignorePath)
		err = ioutil.WriteFile(path.Join(projectRoot, ".gitignore"), []byte(gitignore), os.ModePerm)
		if err != nil {
			util.LogMessageAndExit("Unable to copy .gitignore")
		}
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
