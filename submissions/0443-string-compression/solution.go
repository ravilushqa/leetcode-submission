func compress(chars []byte) int {
    readIdx, writeIdx := 0, 0
    
    for readIdx < len(chars) {
        char := chars[readIdx]
        count := 1
        readIdx++
        
        for readIdx < len(chars) && chars[readIdx] == char {
            count++
            readIdx++
        }

        chars[writeIdx] = char
        writeIdx++
        if count > 1 {
            digits := strconv.Itoa(count)
            for _, c := range []byte(digits) {
                chars[writeIdx] = c
                writeIdx++
            }
        }
    }

    return writeIdx
}
