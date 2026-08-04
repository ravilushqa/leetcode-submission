func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := make(map[byte]int)
	res := 0
	leftIdx := 0

	for rightIdx := 0; rightIdx < len(s); rightIdx++ {
		m[s[rightIdx]]++
		for len(m) > k {
			m[s[leftIdx]]--
			if m[s[leftIdx]] == 0 {
				delete(m, s[leftIdx])
			}
			leftIdx++
		}

		if res < rightIdx-leftIdx+1 {
			res = rightIdx - leftIdx + 1
		}
	}

	return res
}
