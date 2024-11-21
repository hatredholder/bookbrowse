package utils

import (
	"strings"
)

func Truncate(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:strings.LastIndex(s[:max], " ")] + "..."
}

func Commify(slice []string) string {
	return strings.Join(slice, ", ")
}
