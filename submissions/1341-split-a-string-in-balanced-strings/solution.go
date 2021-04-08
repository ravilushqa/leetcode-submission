func balancedStringSplit(s string) int {
    var stack int
	var out int
	for i := range s {
		if s[i] == 'R' {
			stack++
		} else {
			stack--
		}
		if stack == 0 {
			out++
		}
	}
	return out
}
