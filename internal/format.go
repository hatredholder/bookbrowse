package internal

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/utils"
	"github.com/spf13/pflag"
)

func Process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		log.Fatal(err)
	}
	return tmplBytes.String()
}

func ProcessFile(filePath string, vars interface{}, funcMap template.FuncMap) string {
	tmpl, err := template.New(filepath.Base(filePath)).Funcs(funcMap).ParseFiles(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return Process(tmpl, vars)
}

func GetTemplateFile(templateName string) string {
	configDir := utils.GetConfigDir()

	templateFile := filepath.Join(utils.GetTeplatesDir(), templateName+".tmpl")
	if _, err := os.Stat(templateFile); err != nil {
		// No user template found
		templateFile = filepath.Join(configDir, "templates", "predefined", templateName+".tmpl")
		if _, err := os.Stat(templateFile); err != nil {
			fmt.Println("Failed to find template with name:", templateName)
			os.Exit(0)
		}
	}

	return templateFile
}

func Format(book api.Document, flags *pflag.FlagSet) string {
	templateName, err := flags.GetString("template")
	if err != nil {
		log.Fatal(err)
	}
	templateFilePath := GetTemplateFile(templateName)

	funcMap := template.FuncMap{
		"commify":       utils.Commify,
		"truncate":      utils.Truncate,
		"format_rating": utils.FormatRating,
	}

	return ProcessFile(templateFilePath, book, funcMap)
}
