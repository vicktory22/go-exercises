package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var applicationMap = map[rune]string{
	10071:  "recommendation",
	128269: "search",
	9728:   "weather",
}

func Application(log string) string {
	for _, char := range log {
		if application, exists := applicationMap[char]; exists {
			return application
		}
	}

	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	oldString := fmt.Sprintf("%c", oldRune)
	newString := fmt.Sprintf("%c", newRune)

	return strings.ReplaceAll(log, oldString, newString)
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
