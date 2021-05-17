func restoreString(s string, indices []int) string {
    b := make([]byte, len(s))

    for i, v := range indices {
        b[v] = s[i]    
    }
    
    return string(b)
}
