func lengthOfLongestSubstring(s string) int {
	var last [256]int // last[c] = индекс последнего вхождения + 1; 0 = не встречался
	best, left := 0, 0
	for i := 0; i < len(s); i++ {
        c := s[i]
        if last[c] > left {
            left = last[c]
        }

        last[c] = i+1
        if n := i - left + 1;n > best {
            best = n
        }
	}
	return best
}
