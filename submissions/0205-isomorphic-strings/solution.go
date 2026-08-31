func isIsomorphic(s string, t string) bool {
    var sToT,tToS [256]byte

    for i := 0; i < len(s); i++ {
        sVal, tVal := s[i]+1, t[i]+1

        if (sToT[sVal] != 0 && sToT[sVal] != tVal) || (tToS[tVal] != 0 && tToS[tVal] != sVal) {
            return false
        }

        sToT[sVal], tToS[tVal] = tVal, sVal
    }

    return true
}
