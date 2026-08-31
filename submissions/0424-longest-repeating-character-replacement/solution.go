func characterReplacement(s string, k int) int {
    start := 0
    res := 0
    maxFreq := 0
    chars := [26]int{}
    for end, v := range []byte(s) {
        chars[v-'A']++
        windowLen := end - start + 1
        maxFreq = max(maxFreq, chars[v-'A'])

        if windowLen - maxFreq > k {
            chars[s[start]-'A']--
            start++
            windowLen--
        }

        res = max(res, windowLen)
    }

    return res
}
