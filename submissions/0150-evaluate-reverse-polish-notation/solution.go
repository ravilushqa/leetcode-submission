func evalRPN(tokens []string) int {
    stack := []int{}

    for _, v := range tokens{
        switch v {
        case "+",  "-", "/", "*":
            n2 := stack[len(stack) - 1]
            n1 := stack[len(stack) - 2]
            stack = stack[:len(stack) - 2]
            stack = append(stack, execCalc(v, n1, n2))
        default:
            number, err := strconv.Atoi(v)
            if err != nil {
                panic(err)
            }

            stack = append(stack, number)
        }
    }

    return stack[0]
}

func execCalc (token string, n1, n2 int) int {
    switch token {
    case "+":
        return n1+n2
    case "-":
        return n1-n2
    case "/":
        return n1/n2
    case "*":
        return n1*n2
    } 

    panic("unexpected token " + token)
}

