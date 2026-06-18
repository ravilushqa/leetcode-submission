package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var max int
	var sb strings.Builder
	if len(word1) < len(word2) {
		max = len(word2)
	} else {
		max = len(word1)
	}

	for i := 0; i < max; i++ {
		if len(word1) > i {
			sb.WriteByte(word1[i])
		}

		if len(word2) > i {
			sb.WriteByte(word2[i])
		}
	}
	
	return sb.String()
}

