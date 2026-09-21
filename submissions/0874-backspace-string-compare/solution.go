func backspaceCompare(s string, t string) bool {
    stackS := []byte{}
    stackT := []byte{}

    for _, v := range []byte(s) {
        if v == '#' {
            if len(stackS) > 0 {
                stackS = stackS[:len(stackS) - 1]
            }

            continue
        }

        stackS = append(stackS, v)
    }

    for _, v := range []byte(t) {
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
