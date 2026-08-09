func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    var res [][]int

    var i,j int
    for i < len(firstList) && j < len(secondList) {
        a1, a2 := firstList[i][0], firstList[i][1]
        b1, b2 := secondList[j][0], secondList[j][1]

        start, end := max(a1,b1), min(a2,b2)
        if start <= end {
            res = append(res, []int{start, end})
        }

        if end == a2 {
            i++
        } else {
            j++
        }
    }

    return res
}
