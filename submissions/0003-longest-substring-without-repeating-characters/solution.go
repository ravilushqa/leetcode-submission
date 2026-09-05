func lengthOfLongestSubstring(s string) int {
    var start, res int

    chars := map[byte]int{}
    
    for end, char := range []byte(s) {
        chars[char]++
        
        for chars[char] > 1 {
            chars[s[start]]--
            start++
        }

        res = max(res, end - start + 1)
    }

    return res
}
