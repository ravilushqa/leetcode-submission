func isIsomorphic(s string, t string) bool {
    var StoT, TtoS [256]byte

    for i := 0; i < len(s); i++ {
        sVal := s[i]
        tVal := t[i]
        
        if StoT[sVal] == 0 && TtoS[tVal] == 0 {
            StoT[sVal], TtoS[tVal] = tVal, sVal
            continue
        }

        if StoT[sVal] != tVal || TtoS[tVal] != sVal {
            return false
        }
    }

    return true
}
