package words

import "strings"

// CountWords count how many words within a string.
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
