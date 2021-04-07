func numJewelsInStones(J string, S string) int {
    res := 0
	jMap := make(map[rune]struct{}, len(J))
	for _, v := range J{
		jMap[v] = struct{}{}
	}
	for _, char := range S {
		if _, found := jMap[char]; found {
			res++
		}
	}

	return res
}
