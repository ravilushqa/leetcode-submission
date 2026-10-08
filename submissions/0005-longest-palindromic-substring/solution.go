func longestPalindrome(s string) string {
    var max string

    for i := range []byte(s) {
        //odd
        l, r := i, i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if len(max) < r - l + 1 {
                max = s[l:r+1]
            }
            l--
            r++
        }

        //even
        l, r = i, i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if len(max) < r - l + 1 {
                max = s[l:r+1]
            }
            l--
            r++
        }
    }    

    return max 
}

