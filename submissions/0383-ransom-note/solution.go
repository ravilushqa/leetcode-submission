func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int, len(magazine))
    
    for _, v := range magazine {
        m[v]++
    }
    
    for _, v := range ransomNote {
        if m[v] < 1 {
            return false
        }
        m[v]--
    }
    
    return true
}
