func groupAnagrams(strs []string) [][]string {
    groups := map[[26]int][]string{}

    for _, str := range strs {
        chars := [26]int{}
        for _, char := range []byte(str) {
            chars[char-'a']++
        }

        groups[chars] = append(groups[chars], str)
    } 


    res := make([][]string, 0, len(groups))
    for _, group := range groups {
        res = append(res, group)
    }

    return res
}
