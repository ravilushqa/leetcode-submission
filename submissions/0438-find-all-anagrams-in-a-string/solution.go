func findAnagrams(s string, p string) []int {
    var res []int
    var need, window [26]int

    if len(s) < len(p) {
        return res
    }

    for i, v := range p {
        need[v - 'a']++
        window[s[i] - 'a']++
    }

    if need == window {
        res = append(res, 0)
    }

    for i := len(p); i<len(s);i++ {
        window[s[i]-'a']++
        window[s[i-len(p)]-'a']--

        if need == window {
            res = append(res, i-len(p)+1)
        }
    }

    return res
}
