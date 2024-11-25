package internal

import (
	"fmt"

	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/utils"
	"github.com/spf13/pflag"
)

func Format(book api.Document, flags *pflag.FlagSet) string {
	isMarkdown, _ := flags.GetBool("markdown")
	isFullDesc, _ := flags.GetBool("fulldesc")

	title := book.Title
	year := book.Year
	rating := fmt.Sprintf("%.1f", book.Rating*2)
	authors := utils.Commify(book.Authors)
	genres := utils.Commify(book.Genres)
	pages := book.Pages
	image := book.Image.URL
	description := utils.Truncate(book.Description, 350)

	if isFullDesc {
		description = book.Description
	}

	bookInfo := fmt.Sprintf(`
%s (%d) on hardcover.app:

⭐ %s

Author: . . . %s
Genres: . . . %s
Pages:  . . . %d
Plot: . . . . %s
`,
		title, year, rating, authors, genres, pages, description)

	if isMarkdown {
		bookInfo = fmt.Sprintf(`---
genre: %s
pages: %d
rating: %s
my_rating: 
date_started: 
date_finished:
owned: false
---

![image](%s)

# %s, %d

%s

## Authored By

- %s
`,
			genres, pages, rating, image, title, year, description, authors)
	}

	return fmt.Sprint(bookInfo)
}
