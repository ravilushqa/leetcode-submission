func longestPalindrome(s string) string {
    var res string
    for i := range []byte(s) {
        start, end := i, i

        for start >= 0 && end < len(s) && s[start] == s[end] {
            if len(res) < end - start + 1 {
                res = s[start:end+1]
            }
            start--
            end++
        }

        start, end = i, i + 1
        for start >= 0 && end < len(s) && s[start] == s[end] {
            if len(res) < end - start + 1 {
                res = s[start:end+1]
            }
            start--
            end++
        }
    }

    return res
}
