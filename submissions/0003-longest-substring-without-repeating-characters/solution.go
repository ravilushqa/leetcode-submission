func lengthOfLongestSubstring(s string) int {
    res := 0
    start := 0
    chars := [256]int{}

    for end, v := range []byte(s) {
        chars[v]++
        for chars[v] > 1 {
            chars[s[start]]--
            start++
        }

        res = max(res, end - start + 1)
    }

    return res
}
