func compress(chars []byte) int {
    readIdx, writeIdx := 0,  0

    for readIdx < len(chars) {
        ch := chars[readIdx]
        count := 0

        for readIdx < len(chars) && chars[readIdx] == ch {
            count++
            readIdx++
        }

        if count == 1 {
            chars[writeIdx] = ch
            writeIdx++

            continue
        } 

        // more then one
        chars[writeIdx] = ch
        writeIdx++

        for _, digit := range []byte(strconv.Itoa(count)) {
            chars[writeIdx] = digit
            writeIdx++
        }
    }

    return writeIdx
}
