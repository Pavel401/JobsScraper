package utility

import (
	"regexp"
	"strings"
)

// /Using this to pick first three words from a string which starts with capital letter
// /input := "Staff Program Manager Staff Web Engineer Staff Web Engineer Senior Associate, Web Publishing"
// output :=Staff Program Manager
func ExtractThreeWords(input string) string {
	words := strings.Fields(input)
	var result []string

	// Regular expression to match words starting with a capital letter
	re := regexp.MustCompile(`[A-Z][a-zA-Z]*`)

	for _, word := range words {
		if re.MatchString(word) {
			result = append(result, word)
		}
		if len(result) == 3 {
			break
		}
	}

	return strings.Join(result, " ")
}
