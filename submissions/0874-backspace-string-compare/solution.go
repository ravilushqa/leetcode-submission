func backspaceCompare(s string, t string) bool {
    stackS := make([]rune, 0, len(s))
    stackT := make([]rune, 0, len(t))

    for _, v := range s {
        if v == '#' {
            if len(stackS) > 0 {
                stackS = stackS[:len(stackS) - 1]
            }
            continue
        }

        stackS = append(stackS, v)
    }

    for _, v := range t {
        if v == '#' {
            if len(stackT) > 0 {
                stackT = stackT[:len(stackT) - 1]
            }
            continue
        }

        stackT = append(stackT, v)
    }

    return slices.Equal(stackS, stackT)
}
