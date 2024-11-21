package internal

import (
	"fmt"

	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/utils"
)

func Format(book api.Document) string {
	// TODO: return markdown when -m flag is used
	// TODO: return full description when -d flag is used
	// TODO: return /10 rating -r flag is used

	bookInfo := fmt.Sprintf(`
%s (%d) on hardcover.app:

⭐ %s

Author: . . . %s
Genres: . . . %s
Pages:  . . . %d
Plot: . . . . %s
`,
		book.Title, book.Year, fmt.Sprintf("%.1f", book.Rating*2), utils.Commify(book.Authors), utils.Commify(book.Genres), book.Pages, book.Description)

	return fmt.Sprint(bookInfo)
}
