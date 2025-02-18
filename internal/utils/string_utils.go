package utils

import "strings"

// ReplaceSpaces replaces spaces with dashes in the given string.
func ReplaceSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}
