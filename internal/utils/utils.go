package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetConfigDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(userConfigDir, "bookbrowse")
}
