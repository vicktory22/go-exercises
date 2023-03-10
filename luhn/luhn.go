package luhn

import (
	"unicode"
)

func isMinLength(digits []int) bool {
	return len(digits) >= 2
}

func toDigits(id string) []int {
	digits := []int{}

	for _, r := range id {
		if unicode.IsSpace(r) {
			continue
		}

		if !unicode.IsDigit(r) {
			return []int{}
		}

		digits = append(digits, int(r-48))
	}

	return digits
}

func shouldDouble(idx int) bool {
	return idx%2 == 0
}

func sumDigits(digits []int) int {
	sum := 0

	for _, v := range digits {
		sum += v
	}

	return sum
}

func transformDigits(digits []int) []int {
	transformedDigits := []int{}

	if len(digits)%2 != 0 {
		transformedDigits = append(transformedDigits, digits[0])
		digits = digits[1:]
	}

	for i, v := range digits {
		if !shouldDouble(i) {
			transformedDigits = append(transformedDigits, v)
			continue
		}

		doubled := v * 2

		if doubled > 9 {
			doubled -= 9
		}

		transformedDigits = append(transformedDigits, doubled)
	}

	return transformedDigits
}

func Valid(id string) bool {
	digits := toDigits(id)

	if !isMinLength(digits) {
		return false
	}

	transformedDigits := transformDigits(digits)

	return sumDigits(transformedDigits)%10 == 0
}
