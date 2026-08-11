func isIsomorphic(s string, t string) bool {
    var sToT, tToS [256]byte
    for i := 0; i < len(s); i++ {
         a, b := s[i], t[i]
        if sToT[a] != 0 && sToT[a] != b {
            return false
        }
        if tToS[b] != 0 && tToS[b] != a {
            return false
        }
        sToT[a], tToS[b] = b, a
    }


    return true
}
