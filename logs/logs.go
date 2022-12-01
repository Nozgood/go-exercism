package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Application(log string) string {
	var applicationName string 
	for _, char := range log {
		charUnicode := fmt.Sprintf("%U", char)
		if charUnicode == "U+2757" {
			applicationName = "recommendation"
			break
		} else if charUnicode == "U+1F50D" {
			applicationName = "search"
			break
		} else if charUnicode == "U+2600" {
			applicationName = "weather"
		}
	}
	return applicationName
}

func Replace(log string, oldRune, newRune rune) string {
	var valueToReplace string 
	howMuchReplace := 0

	for _, char := range log {
		if char == oldRune {
			valueToReplace = string(char)
			howMuchReplace ++
		}
	}
	newLog := strings.Replace(log, valueToReplace, string(newRune),howMuchReplace)
	return newLog
}

func WithinLimit(log string, limit int) bool {
	logNumberRunes := utf8.RuneCountInString(log)
	return logNumberRunes <= limit
}