package utils

import (
	"regexp"
	"strings"
)

func Slugify(text string) string {
	text = strings.ToLower(text)
	// Replace spaces and special chars with hyphens
	re := regexp.MustCompile(`[^a-z0-9]+`)
	text = re.ReplaceAllString(text, "-")
	text = strings.Trim(text, "-")
	return text
}
