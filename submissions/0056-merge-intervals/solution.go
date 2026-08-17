func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    res := [][]int{}

    for i := range intervals {
        if i == 0 || !isIntersect(res[len(res) - 1], intervals[i]) {
            res = append(res, intervals[i])
        } else {
            res[len(res) - 1][1] = max(res[len(res) - 1][1],intervals[i][1])
        }
    }

    return res
}

func isIntersect(x,y []int) bool {
    return x[1] >= y[0]
}
