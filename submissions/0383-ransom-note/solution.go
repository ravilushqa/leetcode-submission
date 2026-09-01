func canConstruct(ransomNote string, magazine string) bool {
    magazineChars := [26]int{}

    for _, v := range []byte(magazine) {
        magazineChars[v-'a']++
    }

    for _, v := range []byte(ransomNote) {
        magazineChars[v-'a']--
        if magazineChars[v-'a'] < 0 {
            return false
        }
    }

    return true
}
