package main

import (
	"log"

	"github.com/hatredholder/mediabrowse/cmd"
	"github.com/hatredholder/mediabrowse/internal/utils"
)

func main() {
	cmd.Execute()
}

func init() {
	if err := utils.InitConfigDir(); err != nil {
		log.Fatal(err)
	}
}
