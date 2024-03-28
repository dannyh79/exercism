package parsinglogfiles

import (
	"fmt"
	"regexp"
)


func IsValidLine(text string) bool {
	if len(text) < 5 {
		return false
	}

	re := regexp.MustCompile(`^(\[ERR\]|\[INF\])`)
	if re.MatchString(text) {
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~|\*|=|-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i)(password).*"`)
	c := 0
	for _, l := range lines {
		if re.MatchString(l) {
			c += 1
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`^(\[\w+\]).+User\s+(\S+)\s+`)
	for i, l := range lines {
		if match := re.FindStringSubmatch(l); match != nil {
			lines[i] = fmt.Sprintf("[USR] %s %s", match[2], l)
		}
	}
	return lines
}
