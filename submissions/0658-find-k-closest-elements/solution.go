func findClosestElements(arr []int, k int, x int) []int {
    l, r := 0, len(arr) - k

    for l < r {
        m := (l + r) / 2
        
        if IsACloserToX(arr[m], arr[m+k], x) {
            r = m
        } else {
            l = m + 1
        }
    }

    return arr[l:l+k]
}

func IsACloserToX(a, b, x int) bool {
    if a == b {
        return x <= a
    }
    return abs(a-x) < abs(b-x) || (abs(a-x) == abs(b-x) && a < b)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}
