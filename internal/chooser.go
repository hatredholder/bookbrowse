package internal

import (
	"fmt"
	"text/template"

	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/utils"
	"github.com/manifoldco/promptui"
)

func Chooser(hits api.Hits) api.Document {
	funcMap := template.FuncMap{
		"commify": utils.Commify,
		"green":   promptui.Styler(promptui.FGGreen),
		"faint":   promptui.Styler(promptui.FGFaint),
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ .Title }}?",
		Active:   "{{ .Title | green }} ({{.Year}}) by {{ .Authors | commify }}",
		Inactive: "{{ .Title }} ({{.Year}}) by {{ .Authors | commify }}",
		FuncMap:  funcMap,
	}

	prompt := promptui.Select{
		Label:        "Select Book",
		Items:        hits,
		Templates:    templates,
		HideHelp:     true,
		HideSelected: true,
	}

	// TODO: do error checking
	i, _, err := prompt.Run()
	// _, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return hits[i].Document
}
