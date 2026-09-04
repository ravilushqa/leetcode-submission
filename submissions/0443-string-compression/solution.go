func compress(chars []byte) int {
    var readIdx, writeIdx int

    for readIdx < len(chars) {
        char := chars[readIdx]
        count := 0
        
        for readIdx < len(chars) && chars[readIdx] == char {
            count++
            readIdx++
        }

        chars[writeIdx] = char
        writeIdx++

        if count == 1 {
            continue
        }

        digits := strconv.Itoa(count)
        for _, digit := range []byte(digits) {
            chars[writeIdx] = digit
            writeIdx++
        }
    }

    return writeIdx
}
