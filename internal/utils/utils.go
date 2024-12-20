package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hatredholder/mediabrowse/internal/templates"
)

func GetConfigDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(userConfigDir, "mediabrowse")
}

func FindAvailableTmpls() string {
	availableTmpls, err := os.ReadDir(GetConfigDir())
	if err != nil {
		log.Fatal(err)
	}

	result := []string{}
	for _, f := range availableTmpls {
		if strings.HasSuffix(f.Name(), ".tmpl") {
			withsuffix := f.Name()
			suffixless := strings.TrimSuffix(withsuffix, filepath.Ext(withsuffix))
			result = append(result, suffixless)
		}
	}
	return strings.Join(result, ", ")
}

func CreateTmplFiles() error {
	if err := os.WriteFile(GetConfigDir()+"/default.tmpl", []byte(templates.DefaultTmpl), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(GetConfigDir()+"/markdown.tmpl", []byte(templates.MarkdownTmpl), 0644); err != nil {
		return err
	}
	return nil
}

func InitConfigDir() error {
	if _, err := os.Stat(GetConfigDir()); os.IsNotExist(err) {
		if err := os.MkdirAll(GetConfigDir(), os.ModePerm); err != nil {
			return err
		}
		if err := CreateTmplFiles(); err != nil {
			return err
		}
	}
	return nil
}
