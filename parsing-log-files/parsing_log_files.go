package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var isValidLogRegex = regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
var splitLogRegex = regexp.MustCompile(`\<[~|\*|=|-]*\>`)
var hasPasswordRegex = regexp.MustCompile(`(?i)".*password.*"`)
var endOfLineRegex = regexp.MustCompile(`end-of-line[\d]+`)
var userNameRegex = regexp.MustCompile(`(?-i)User\s+(?P<username>.[^\s]*)`)

func IsValidLine(text string) bool {
	return isValidLogRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	total := 0

	for _, line := range lines {
		if hasPasswordRegex.MatchString(line) {
			total++
		}
	}

	return total
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for idx, line := range lines {
		match := userNameRegex.FindStringSubmatch(line)

		if len(match) == 0 {
			continue
		}

		lines[idx] = fmt.Sprintf("[USR] %s %s", match[1], line)
	}

	return lines
}
