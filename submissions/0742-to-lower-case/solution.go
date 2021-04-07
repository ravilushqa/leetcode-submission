func toLowerCase(str string) string {
	var res []rune
	for _, char := range str {
		if char > 64 && char < 91 {
			res = append(res, char+32)
		} else {
			res = append(res, char)
		}
	}

	return string(res)
}

