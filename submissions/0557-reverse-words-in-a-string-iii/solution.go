func reverseWords(s string) string {
    words := strings.Split(s, " ")
	for i, word := range words {
		runes := []rune(word)
		for i := 0; i <= (len(word)-1)/2; i++ {
			runes[i], runes[len(runes) - 1 - i] = runes[len(runes) - 1 - i], runes[i]
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}
