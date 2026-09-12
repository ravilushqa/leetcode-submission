func characterReplacement(s string, k int) int {
    start := 0
    res := 0
    chars := [26]int{}
    maxFreq := 0

    for end, v := range s {
        chars[v-'A']++
        
        maxFreq = max(maxFreq, chars[v-'A'])
        windowLen := end - start + 1

        for windowLen - maxFreq > k {
            windowLen--
            chars[s[start]-'A']--
            start++
        }

        res = max(res, windowLen)
    }

    return res
}
