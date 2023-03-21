package leap

func IsLeapYear(year int) bool {
	isLeapYear := (year % 4) == 0

	if !isLeapYear {
		return false
	}

	isYear100Exception := (year % 100) == 0
	isYear400Exception := (year % 400) == 0

	if isYear100Exception && !isYear400Exception {
		return false
	}

	return true
}
