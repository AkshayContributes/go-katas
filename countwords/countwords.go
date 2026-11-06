package countwords

import "strings"

func CountWordsInString(s string) int {
	return len(strings.Fields(s))
}

func CountWordFrequency(s string) map[string]int {
	freq := make(map[string]int)

	for _, word := range strings.Fields(s) {
		if strings.TrimSpace(word) != "" {
			freq[word]++
		}
	}

	return freq
}
