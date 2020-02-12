package util

import (
	"bytes"
	"text/template"
)

func ParseTemplate(templateData interface{}, templateFile, templateStr string) (string, error) {
	t, err := template.New(templateFile).Parse(templateStr)
	if err != nil {
		return "", nil
	}
	var tmpl bytes.Buffer
	err = t.ExecuteTemplate(&tmpl, templateFile, templateData)
	if err != nil {
		return "", nil
	}
	return tmpl.String(), nil
}
