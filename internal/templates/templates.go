package templates

var (
	DefaultTmpl  string
	MarkdownTmpl string
)

func init() {
	DefaultTmpl = `{{ .Title }} ({{ .Year }}) on hardcover.app:

⭐ {{ .Rating | formatRating }}

Author: . . . {{ .AuthorNames | commify }}
Genres: . . . {{ .Genres | commify }}
Pages:  . . . {{ .Pages }}
Plot: . . . . {{ truncate .Description 350 }}
`
	MarkdownTmpl = `---
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
