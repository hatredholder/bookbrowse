package internal

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/templates"
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

func ProcessFile(tmplPath string, funcMap template.FuncMap, vars interface{}) string {
	tmplFile := filepath.Base(tmplPath)

	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	return Process(tmpl, vars)
}

func GetTemplateFile(templateName string) string {
	templateFile := filepath.Join(utils.GetConfigDir(), templateName+".tmpl")
	if _, err := os.Stat(templateFile); err != nil {
		fmt.Println("Failed to find template with name:", templateName)
		os.Exit(0)
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
		"commify":      templates.Commify,
		"truncate":     templates.Truncate,
		"formatRating": templates.FormatRating,
	}

	return ProcessFile(templateFilePath, funcMap, book)
}
