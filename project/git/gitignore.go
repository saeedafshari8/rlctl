package git

import (
	"github.com/rocketlaunchercloud/rlctl/util"
	"io/ioutil"
	"os"
	"path"
)

var (
	gitignorePath = "buildpipeline/.gitignore.tmpl"
)

func ParseAndSaveGitIgnore(projectRootPath string) error {
	gitignore, err := util.GetSpringTemplate(gitignorePath)
	err = ioutil.WriteFile(path.Join(projectRootPath, ".gitignore"), []byte(gitignore), os.ModePerm)
	return err
}
