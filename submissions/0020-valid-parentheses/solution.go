func isValid(s string) bool {
    stack := []rune{}

    pairs := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    for _, v := range []rune(s) {
        switch v {
        case '(', '{', '[':
            stack = append(stack, v)
        case ')', '}', ']':
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            
            if pair, found := pairs[top]; found && pair == v {
                stack = stack[:len(stack)-1]
                continue
            }

            return false
        }
    }

    return len(stack) == 0
}
