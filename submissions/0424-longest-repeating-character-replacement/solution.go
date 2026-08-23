func characterReplacement(s string, k int) int {
    res := 0
    start := 0

    var freqs [26]int

    for end, v := range []byte(s) {
        freqs[v-'A']++
        maxFreq := GetMaxFreq(freqs)
        windowLen := end - start + 1
        for windowLen - maxFreq > k {
            freqs[s[start] - 'A']--
            start++
            windowLen--
            maxFreq = GetMaxFreq(freqs)
        }

        res = max(res, windowLen)
    }

    return res
}

func GetMaxFreq(freqs [26]int) int {
    var res int
    for _,v := range freqs {
        res = max(res, v)
    }

    return res
}
