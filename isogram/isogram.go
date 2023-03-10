package isogram

import (
	"unicode"
)

func isInvalidChar(r rune) bool {
	return r == ' ' || r == '-'
}

func IsIsogram(word string) bool {
	charMap := make(map[rune]bool)

	for _, char := range word {
		uchar := unicode.ToUpper(char)

		if isInvalidChar(char) {
			continue
		}

		if charMap[uchar] == true {
			return false
		}

		charMap[uchar] = true
	}

	return true
}
