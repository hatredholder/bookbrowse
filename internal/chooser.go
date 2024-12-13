package internal

import (
	"fmt"
	"os"
	"text/template"

	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/templates"
	"github.com/manifoldco/promptui"
)

func Chooser(hits api.Hits) api.Document {
	if len(hits) == 0 {
		fmt.Println("No results found")
		os.Exit(0)
	}

	funcMap := template.FuncMap{
		"commify": templates.Commify,
		"green":   promptui.Styler(promptui.FGGreen),
		"faint":   promptui.Styler(promptui.FGFaint),
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ .Title }}?",
		Active:   "{{ .Title | green }} ({{ .Year }}) by {{ .AuthorNames | commify }}",
		Inactive: "{{ .Title }} ({{ .Year }}) by {{ .AuthorNames | commify }}",
		FuncMap:  funcMap,
	}

	prompt := promptui.Select{
		Label:        "Select Book",
		Items:        hits,
		Templates:    templates,
		HideHelp:     true,
		HideSelected: true,
		Stdout:       os.Stderr,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		os.Exit(1)
	}

	return hits[i].Document
}
