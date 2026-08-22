func characterReplacement(s string, k int) int {
    res := 0
    start := 0

    var chars [26]int
    maxFreq := 0
    for end, v := range []byte(s) {
        chars[v - 'A']++

        windowLen := end - start + 1
        maxFreq = max(maxFreq, chars[v - 'A'])

        for windowLen - maxFreq > k{
            chars[s[start] - 'A']--
            start++
            windowLen--
        }

        res = max(res, end - start + 1)
    }

    return res
}
