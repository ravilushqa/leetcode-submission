// a-z
// bruteforce: n^2
// s: abababa 3+2+1 = 6
// abab
func numberOfSpecialSubstrings(s string) int {
    res := 0
    start := 0

    var chars [26]int

    for end, char := range s {
        chars[char-'a']+=1

        for chars[char-'a'] > 1 {
            chars[s[start]-'a']--
            start++
        }
        
        res += end - start + 1
    } 

    return res

}
