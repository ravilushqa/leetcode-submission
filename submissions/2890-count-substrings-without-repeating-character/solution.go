func numberOfSpecialSubstrings(s string) int {
	start := 0
	res := 0

	chars := [26]int{}
	for end, v := range []byte(s) {
		chars[v-'a']++
		for chars[v-'a'] > 1 {
			chars[s[start]-'a']--
			start++
		}

		res += end - start + 1
	}

	return res
}
