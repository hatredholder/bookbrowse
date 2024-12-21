package internal

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/hatredholder/mediabrowse/internal/api"
	"github.com/hatredholder/mediabrowse/internal/templates"
	"github.com/hatredholder/mediabrowse/internal/utils"
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

func GetTmplFile(tmplName string) string {
	tmplFile := filepath.Join(utils.GetConfigDir(), tmplName+".tmpl")
	if _, err := os.Stat(tmplFile); err != nil {
		fmt.Println("Failed to find template with name:", tmplName)
		fmt.Println("Available templates:", utils.FindAvailableTmpls())
		os.Exit(0)
	}

	return tmplFile
}

func Format(book api.Document, flags *pflag.FlagSet) string {
	tmplName, err := flags.GetString("template")
	if err != nil {
		log.Fatal(err)
	}

	tmplPath := GetTmplFile(tmplName)

	funcMap := template.FuncMap{
		"commify":      templates.Commify,
		"truncate":     templates.Truncate,
		"formatRating": templates.FormatRating,
	}

	return ProcessFile(tmplPath, funcMap, book)
}
