func longestPalindrome(s string) string {
	res := s[0:1]

	for i := 0; i <= len(s)-1; i++ { //i 0 -> len-1
		for j := len(s); j > i; j-- { //j len ->i
			if len(res) >= j-i {
				continue
			}
			if isPalindrome(s[i:j]) {
				res = s[i:j]
			}
		}
	}

	return res
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
