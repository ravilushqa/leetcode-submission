func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	runes := []rune(strconv.Itoa(x))

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
