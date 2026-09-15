func backspaceCompare(s string, t string) bool {
    stackS := []byte{}
    stackT := []byte{}

    for _, v := range []byte(s) {
        if v != '#' {
            stackS = append(stackS, v)

            continue
        }

        if len(stackS) > 0 {
            stackS = stackS[:len(stackS) - 1]
        }
    }

    for _, v := range []byte(t) {
        if v != '#' {
            stackT = append(stackT, v)

            continue
        }

        if len(stackT) > 0 {
            stackT = stackT[:len(stackT) - 1]
        }
    }

    return slices.Equal(stackS, stackT)
}
