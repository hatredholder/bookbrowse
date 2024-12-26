package internal

import (
	"bytes"
	"log"
	"path/filepath"
	"text/template"

	"github.com/hatredholder/mediabrowse/internal/api"
	"github.com/hatredholder/mediabrowse/internal/templates"
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

func Format(book api.Document, tmplPath string) string {
	funcMap := template.FuncMap{
		"commify":      templates.Commify,
		"truncate":     templates.Truncate,
		"formatRating": templates.FormatRating,
	}

	return ProcessFile(tmplPath, funcMap, book)
}
