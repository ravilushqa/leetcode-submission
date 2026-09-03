func longestPalindrome(s string) int {
    var res int

    chars := make(map[byte]int, 52)

    for _, v := range []byte(s) {
        chars[v]++

        if chars[v] == 2 {
            res +=2

            delete(chars, v)
        }
    }

    if len(chars) > 0 {
        res++
    }

    return res
}
