func firstUniqChar(s string) int {
    chars := [26]int{}

    for _, v := range []byte(s) {
        chars[v-'a']++
    }

    for i, v := range []byte(s) {
        if chars[v-'a'] == 1 {
            return i
        }
    }

    return -1
}
