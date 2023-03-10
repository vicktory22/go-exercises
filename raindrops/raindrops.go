package raindrops

import "strconv"

func isPling(number int) bool {
	return number%3 == 0
}

func isPlang(number int) bool {
	return number%5 == 0
}

func isPlong(number int) bool {
	return number%7 == 0
}

func Convert(number int) string {
	phrase := ""

	if isPling(number) {
		phrase += "Pling"
	}

	if isPlang(number) {
		phrase += "Plang"
	}

	if isPlong(number) {
		phrase += "Plong"
	}

	if phrase != "" {
		return phrase
	}

	return strconv.Itoa(number)
}
