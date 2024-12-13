package templates

import (
	"os"

	"github.com/hatredholder/bookbrowse/internal/utils"
)

var (
	defaultTmpl  string
	markdownTmpl string
)

func init() {
	defaultTmpl = `{{ .Title }} ({{ .Year }}) on hardcover.app:

⭐ {{ .Rating | formatRating }}

Author: . . . {{ .AuthorNames | commify }}
Genres: . . . {{ .Genres | commify }}
Pages:  . . . {{ .Pages }}
Plot: . . . . {{ truncate .Description 350 }}
`
	markdownTmpl = `---
genre: {{ .Genres | commify }}
pages: {{ .Pages }}
rating: {{ .Rating | formatRating }}
my_rating: 
date_started: 
date_finished:
owned: false
---

![image]({{ .Image.URL }})

# {{ .Title }}, {{ .Year }}

{{ .Description }}

## Authored By

- {{ .AuthorNames | commify }}
`
}

func CreateTmplFiles() error {
	if err := os.WriteFile(utils.GetConfigDir()+"/default.tmpl", []byte(defaultTmpl), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(utils.GetConfigDir()+"/markdown.tmpl", []byte(markdownTmpl), 0644); err != nil {
		return err
	}
	return nil
}
