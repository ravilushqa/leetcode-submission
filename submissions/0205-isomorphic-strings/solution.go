func isIsomorphic(s string, t string) bool {
    var tToS, sToT [256]byte

    for i := 0; i < len(s); i++ {
        sVal, tVal := s[i]+1, t[i]+1
        
        if (tToS[tVal] != 0 && tToS[tVal] != sVal) || (sToT[sVal] != 0 && sToT[sVal] != tVal) {
            return false
        }

        tToS[tVal], sToT[sVal] =  sVal, tVal
    } 

    return true
}
