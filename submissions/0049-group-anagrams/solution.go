func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string, len(strs))

    for _, str := range strs {
        var chars [26]int

        for _, char := range str {
            chars[char-'a']++
        }

        m[chars] = append(m[chars], str)
    } 

    res := make([][]string, 0, len(m))
    for _, strs := range m {
        res = append(res, strs)
    }

    return res
}
