func addBinary(a string, b string) string {
    max := max(len(a), len(b))
    res := make([]byte, max + 1)

    aPtr, bPtr := len(a) - 1, len(b) - 1
    carry := 0
    for i := 0; aPtr >= 0 || bPtr >= 0 || carry > 0; i, aPtr, bPtr = i + 1, aPtr - 1, bPtr - 1  {
        sumRes := 0
        sumRes += carry
        if aPtr >= 0 && a[aPtr] == '1' {
            sumRes++
        }
        if bPtr >= 0 && b[bPtr] == '1' {
            sumRes++
        }

        writeIdx := len(res) - 1 - i
        switch sumRes {
        case 0:
            res[writeIdx] = '0'
            carry = 0
        case 1:
            res[writeIdx] = '1'
            carry = 0
        case 2:
            res[writeIdx] = '0'
            carry = 1
        case 3:
            res[writeIdx] = '1'
            carry = 1
        default:
            panic("unexpected res")
        }
    }

    if int(res[0]) == 0 {
        return string(res[1:])
    }
    
    return string(res)
}
