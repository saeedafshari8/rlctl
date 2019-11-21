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

var (
	applicationConfigTemplate            = "project/java/spring/config/application.yml.tmpl"
	applicationLocalConfigTemplate       = "project/java/spring/config/application-local.yml.tmpl"
	applicationIntegrationConfigTemplate = "project/java/spring/config/application-int.yml.tmpl"
	applicationProdConfigTemplate        = "project/java/spring/config/application-prod.yml.tmpl"
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

	compileTemplateAndSave(&configPath, &applicationConfigTemplate, templateData, "application.yml")
	compileTemplateAndSave(&configPath, &applicationLocalConfigTemplate, templateData, "application-local.yml")
	compileTemplateAndSave(&configPath, &applicationIntegrationConfigTemplate, templateData, "application-int.yml")
	compileTemplateAndSave(&configPath, &applicationProdConfigTemplate, templateData, "application-prod.yml")
}

func compileTemplateAndSave(configPath, templatePath *string, templateData *ProjectConfig, fileName string) {
	filePath := path.Join(*configPath, fileName)
	err := ioutil.WriteFile(filePath, []byte(parseApplicationTemplate(templateData, *templatePath)), os.ModePerm)
	if err != nil {
		util.LogMessageAndExit(fmt.Sprintf("Unable to save %s", filePath))
	}
	log.Printf("%s config file created successfully!", fileName)
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
