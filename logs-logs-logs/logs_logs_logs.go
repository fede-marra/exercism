package logs

import (
	"fmt"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		cha := fmt.Sprintf("%U", char)

		if cha == "U+2757" {
			return "recommendation"
		}
		if cha == "U+1F50D" {
			return "search"
		}
		if cha == "U+2600" {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, fmt.Sprintf("%c", oldRune), fmt.Sprintf("%c", newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0
	for range log {
		count++
	}
	return count <= limit
}
