package utils

import (
	"strings"
)

// TODO: create a description shortener function
// func Shorten() {}

func Commify(slice []string) string {
	return strings.Join(slice, ", ")
}
