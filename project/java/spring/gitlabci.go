package spring

import (
	"github.com/saeedafshari8/flixinit/util"
	"os"
	"path"
)

var (
	moPath = "project/java/spring/buildpipeline/mo.sh"
)

func ParseAndSaveCiCdFile(projectRoot string, templateData *ProjectConfig) {
	if (*templateData).EnableGitLab {
		configPath := path.Join(projectRoot, "build_pipeline")
		util.CreateDirIfNotExists(&configPath)

		cwd, err := os.Getwd()
		util.LogAndExit(err, util.EnvironmentError)
		_, err = util.Copy(path.Join(cwd, moPath), path.Join(configPath, "mo.sh"))
		if err != nil {
			util.LogMessageAndExit("Unable to copy Liquibase master.xml")
		}
	}
}
