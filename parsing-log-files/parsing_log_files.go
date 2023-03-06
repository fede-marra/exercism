package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	valid := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}
	for _, val := range valid {
		if text != "" && val == text[1:4] {
			return true
		}
	}
	return false
}
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<(\*|\=|\~|\-)*\>`)
	return re.Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	var cont int
	re := regexp.MustCompile(`(?i)\".*password.*\"`)
	for _, text := range lines {
		if re.MatchString(text) {
			cont++
		}
	}
	return cont
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile("end-of-line[[:digit:]]{0,}")
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var log []string
	re := regexp.MustCompile(`User[[:space:]]+([[:alnum:]]+)[[:space:]]+.*`)
	for _, text := range lines {
		matchingLine := re.FindStringSubmatch(text)
		if len(matchingLine) == 0 {
			log = append(log, text)
		} else {
			log = append(log, fmt.Sprintf("[USR] %s %s", matchingLine[1], text))
		}
	}
	return log
}
