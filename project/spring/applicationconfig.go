package spring

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	applicationConfigTemplate            = "config/application.yml.tmpl"
	applicationLocalConfigTemplate       = "config/application-local.yml.tmpl"
	applicationIntegrationConfigTemplate = "config/application-int.yml.tmpl"
	applicationProdConfigTemplate        = "config/application-prod.yml.tmpl"
	liquibaseConfigTemplate              = "config/liquibase-master.xml.tmpl"
)

func ParseAndSaveAppConfigTemplates(projectRoot string, templateData *SpringProjectConfig) {
	configPath := path.Join(projectRoot, "config")
	util.CreateDirIfNotExists(&configPath)

	if (*templateData).EnableLiquibase {
		liquibaseDbChangeSetPath := path.Join(projectRoot, "src/main/resources/db")
		util.CreateDirIfNotExists(&liquibaseDbChangeSetPath)

		liquibaseTemplate, err := util.GetSpringTemplate(liquibaseConfigTemplate)
		if err != nil {
			util.LogMessageAndExit("Unable to copy Liquibase master.xml")
		}
		liquibaseParsedTemplate, err := util.ParseTemplate(templateData, "master.xml", liquibaseTemplate)
		if err != nil {
			util.LogMessageAndExit("Unable to copy Liquibase master.xml")
		}

		err = ioutil.WriteFile(path.Join(liquibaseDbChangeSetPath, "master.xml"), []byte(liquibaseParsedTemplate), os.ModePerm)
		if err != nil {
			util.LogMessageAndExit("Unable to copy Liquibase master.xml")
		}
	}

	compileTemplateAndSave(&configPath, &applicationConfigTemplate, templateData, "application.yml")
	compileTemplateAndSave(&configPath, &applicationLocalConfigTemplate, templateData, "application-local.yml")
	compileTemplateAndSave(&configPath, &applicationIntegrationConfigTemplate, templateData, "application-int.yml")
	compileTemplateAndSave(&configPath, &applicationProdConfigTemplate, templateData, "application-prod.yml")
}

func compileTemplateAndSave(configPath, templatePath *string, templateData *SpringProjectConfig, fileName string) {
	springTemplate, err := util.GetSpringTemplate(*templatePath)
	util.LogAndExit(err, util.InvalidTemplate)

	parsedTemplate, err := util.ParseTemplate(templateData, fileName, springTemplate)
	util.LogAndExit(err, util.InvalidTemplate)

	filePath := path.Join(*configPath, fileName)
	err = ioutil.WriteFile(filePath, []byte(parsedTemplate), os.ModePerm)
	if err != nil {
		util.LogMessageAndExit(fmt.Sprintf("Unable to save %s", filePath))
	}
	log.Printf("%s config file created successfully!", fileName)
}
