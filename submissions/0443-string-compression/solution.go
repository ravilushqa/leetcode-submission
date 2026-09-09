func compress(chars []byte) int {
    var write, read int

    for read < len(chars) {
        char := chars[read]
        count := 0
        for read < len(chars) && char == chars[read]{
            count++
            read++
        }

        chars[write] = char
        write++

        if count > 1 {
            digits := strconv.Itoa(count)

            for _, v := range []byte(digits) {
                chars[write] = v
                write++
            }
        }
    }

    return write
}
