package internal

import (
	"bytes"
	"log"
	"path/filepath"
	"text/template"

	"github.com/hatredholder/mediabrowse/internal/api"
	"github.com/hatredholder/mediabrowse/internal/templates"
)

func processTmpl(book api.Document, tmpl *template.Template) string {
	var tmplBytes bytes.Buffer

	err := tmpl.Execute(&tmplBytes, book)
	if err != nil {
		log.Fatal(err)
	}

	return tmplBytes.String()
}

func Format(book api.Document, tmplPath string) string {
	tmplFile := filepath.Base(tmplPath)

	funcMap := template.FuncMap{
		"commify":      templates.Commify,
		"truncate":     templates.Truncate,
		"formatRating": templates.FormatRating,
	}

	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	return processTmpl(book, tmpl)
}
