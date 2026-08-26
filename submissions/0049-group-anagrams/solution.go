func groupAnagrams(strs []string) [][]string {
    chars := make(map[[26]int][]string, len(strs))

    for _, str := range strs {
        var strChars [26]int
        for _, char := range []byte(str) {
            strChars[char - 'a']++
        }

        chars[strChars] = append(chars[strChars], str)
    }

    res := make([][]string, 0,len(chars))
    for _, v := range chars {
        res = append(res, v)
    }

    return res
}
