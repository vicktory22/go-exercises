package scrabble

import (
	"unicode"
)

var wordMap = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func getSingleLetterMap(words map[string]int) map[rune]int {
	singleLetterMap := make(map[rune]int)

	for word, score := range words {
		for _, letter := range word {
			singleLetterMap[letter] = score
		}
	}

	return singleLetterMap
}

func Score(word string) int {
	singleLetterMap := getSingleLetterMap(wordMap)

	score := 0

	for _, letter := range word {
		score += singleLetterMap[unicode.ToUpper(letter)]
	}

	return score
}
