func lengthOfLongestSubstring(s string) int {
    var max int

	for i := range s {
		m := make(map[rune]struct{}, len(s))
		iterMax := 0
		for _, v := range s[i:] {
			if _, ok := m[v]; ok {
				break
			}

			m[v] = struct{}{}
			iterMax++
		}

		if max < iterMax {
			max = iterMax
		}
	}

	return max
}
