package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetConfigDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(userConfigDir, "bookbrowse")
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
