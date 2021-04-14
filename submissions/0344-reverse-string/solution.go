func reverseString(s []byte)  {
    for i := range s {
        if i >= len(s)/2 {
            break
        }
        
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
}
