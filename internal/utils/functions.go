package utils

import (
	"fmt"
	"strings"
)

func Commify(slice []string) string {
	return strings.Join(slice, ", ")
}

func Truncate(s string, max int) string {
	if 350 > len(s) {
		return s
	}
	return s[:strings.LastIndex(s[:max], " ")] + "..."
}

func FormatRating(rating float64) string {
	return fmt.Sprintf("%.1f", rating*2)
}
