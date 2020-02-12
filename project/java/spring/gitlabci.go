package spring

import (
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"os"
	"path"
)

var (
	moPath           = "buildpipeline/mo.sh"
	gitignorePath    = "buildpipeline/.gitignore.tmpl"
	gitlabCITemplate = "buildpipeline/.gitlab-ci-default.yml"
)

type GitLabCI struct {
	Tags    []string
	Excepts []string
}

func ParseAndSaveCiCdFile(projectRoot string, templateData *ProjectConfig) {
	if (*templateData).EnableGitLab {
		configPath := path.Join(projectRoot, "build_pipeline")
		util.CreateDirIfNotExists(&configPath)

		mo, err := util.GetSpringTemplate(moPath)
		err = ioutil.WriteFile(path.Join(configPath, "mo.sh"), []byte(mo), os.ModePerm)
		if err != nil {
			util.LogMessageAndExit("Unable to copy mo.sh")
		}

		gitignore, err := util.GetSpringTemplate(gitignorePath)
		err = ioutil.WriteFile(path.Join(projectRoot, ".gitignore"), []byte(gitignore), os.ModePerm)
		if err != nil {
			util.LogMessageAndExit("Unable to copy .gitignore")
		}
	}

	compileTemplateAndSave(&projectRoot, &gitlabCITemplate, templateData, ".gitlab-ci.yml")
}
