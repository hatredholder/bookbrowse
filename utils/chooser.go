package utils

import (
	"github.com/hatredholder/bookbrowse/api"
	"github.com/manifoldco/promptui"
)

func Chooser(hits api.Hits) api.Document {
	// TODO: remove the grey texting appearing after making a selection
	// TODO: use the commify function for authors
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Title }}?",
		Active:   "{{ .Title | cyan }} ({{.Year}}) by {{ .Authors }}",
		Inactive: "{{ .Title }} ({{.Year}}) by {{ .Authors }}",
	}

	prompt := promptui.Select{
		Label:     "Select Book",
		Items:     hits,
		Templates: templates,
		HideHelp:  true,
	}

	// TODO: do error checking
	i, _, _ := prompt.Run()
	// _, result, err := prompt.Run()
	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	return hits[i].Document
}
