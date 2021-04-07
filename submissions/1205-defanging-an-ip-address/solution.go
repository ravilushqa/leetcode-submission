func defangIPaddr(address string) string {
    	var defanged []rune

	for _, c := range address {
		if c == '.' {
			defanged = append(defanged,'[','.',']')
		} else {
			defanged = append(defanged,c)
		}
	}

	return string(defanged)
}
