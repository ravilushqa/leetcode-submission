func numberOfSpecialSubstrings(s string) int {
    res := 0
    start := 0

    var chars [26]int

    for end, v := range s {
        chars[v-'a']++

        for chars[v-'a'] > 1 {
            chars[s[start] - 'a']--
            start++
        }

        res += end - start + 1
    }

    return res
}
