func decodeString(s string) string {
    stack := []byte{}

    for _, v := range []byte(s) {
        if v != ']' {
            stack = append(stack, v)
            continue
        }

        str := ""
        for stack[len(stack) - 1] != '[' {
            str = string(stack[len(stack) - 1]) + str
            stack = stack[:len(stack) - 1]
        }
        stack = stack[:len(stack) - 1]
        
        k := ""
        for len(stack) > 0 && stack[len(stack) - 1] >= '0' && stack[len(stack) - 1] <= '9' {
            k = string(stack[len(stack) - 1]) + k
            stack = stack[:len(stack) - 1]
        }

        kInt, _ := strconv.Atoi(k)


        stack = append(stack, []byte(strings.Repeat(str, kInt))...)
    }

    return string(stack)
}

