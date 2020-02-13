package spring

import (
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"os"
	"path"
)

const (
	K8SProdTemplate    = "kubernetes/prod/kube-config.yml"
	K8SStagingTemplate = "kubernetes/stg/kube-config.yml"
)

func parseK8STemplates(projectConfig *SpringProjectConfig) (string, string) {
	prodTmplStr, err := util.GetSpringTemplate(K8SProdTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	parsedProdTemplate, err := util.ParseTemplate(projectConfig, "kube-config", prodTmplStr)
	util.LogAndExit(err, util.InvalidTemplate)

	stgTmplStr, err := util.GetSpringTemplate(K8SStagingTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	parsedStgTemplate, err := util.ParseTemplate(projectConfig, "kube-config", stgTmplStr)
	util.LogAndExit(err, util.InvalidTemplate)

	return parsedProdTemplate, parsedStgTemplate
}

func SaveK8sTemplates(projectRoot *string, projectConfig *SpringProjectConfig) {
	prod, stg := parseK8STemplates(projectConfig)

	kubernetesConfigPath := path.Join(*projectRoot, "kubernetes")
	util.CreateDirIfNotExists(&kubernetesConfigPath)

	stgPath := path.Join(*projectRoot, "kubernetes/stg")
	util.CreateDirIfNotExists(&stgPath)

	prodPath := path.Join(*projectRoot, "kubernetes/prod")
	util.CreateDirIfNotExists(&prodPath)

	err := ioutil.WriteFile(path.Join(stgPath, "kube-config.yml"), []byte(stg), os.ModePerm)
	util.LogAndExit(err, util.FileNotFound)

	err = ioutil.WriteFile(path.Join(prodPath, "kube-config.yml"), []byte(prod), os.ModePerm)
	util.LogAndExit(err, util.FileNotFound)
}
