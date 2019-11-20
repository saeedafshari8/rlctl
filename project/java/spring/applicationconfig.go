package spring

import (
	"bytes"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

const (
	applicationConfigTemplate = "project/java/spring/config/application.yml.tmpl"
)

func ParseAndSaveAppConfigTemplates(projectRoot string, templateData *ProjectConfig) {
	configPath := path.Join(projectRoot, "config")
	util.CreateDirIfNotExists(&configPath)

	if (*templateData).EnableLiquibase {
		liquibaseDbChangeSetPath := path.Join(projectRoot, "src/main/resources/db")
		util.CreateDirIfNotExists(&liquibaseDbChangeSetPath)
		cwd, err := os.Getwd()
		util.LogAndExit(err, util.EnvironmentError)
		_, err = util.Copy(path.Join(cwd, "project/java/spring/config/liquibase-master.xml.tmpl"),
			path.Join(liquibaseDbChangeSetPath, "master.xml"))
		if err != nil {
			util.LogMessageAndExit("Unable to copy Liquibase master.xml")
		}
	}

	filePath := path.Join(configPath, "application.yml")
	err := ioutil.WriteFile(filePath, []byte(parseApplicationTemplate(templateData, applicationConfigTemplate)), os.ModePerm)
	if err != nil {
		util.LogMessageAndExit(fmt.Sprintf("Unable to save %s", filePath))
	}
	log.Println("Application config files created successfully!")
}

func parseApplicationTemplate(templateData *ProjectConfig, templateFile string) string {
	dir, err := os.Getwd()
	util.LogAndExit(err, util.EnvironmentError)
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, templateFile))
	util.LogAndExit(err, util.FileNotFound)
	t, err := template.New(templateFile).Parse(string(file))
	util.LogAndExit(err, util.InvalidTemplate)
	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, templateFile, *templateData)
	util.LogAndExit(err, util.InvalidTemplate)
	return tmpl.String()
}
