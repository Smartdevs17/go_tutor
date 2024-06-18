// Count how many times each word appears in a text
package sample

import (
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	return counts
}
