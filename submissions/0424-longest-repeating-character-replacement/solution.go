func characterReplacement(s string, k int) int {
    res := 0
    start := 0
    freqs := [26]int{}

    for end, v := range []byte(s) {
        freqs[v-'A']++

        windowLen := end - start + 1
        maxFreq := topFreq(freqs)

        for windowLen - maxFreq > k {
            freqs[s[start]-'A']--
            start++
            windowLen--
            maxFreq = topFreq(freqs)
        }

        res = max(res, windowLen)
    }

    return res
}

func topFreq(freqs [26]int) int {
    res := 0
    for _, v := range freqs {
        res = max(res, v)
    }

    return res
}
