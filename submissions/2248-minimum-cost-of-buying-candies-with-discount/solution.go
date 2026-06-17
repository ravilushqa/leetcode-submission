import (
    "sort"
)

func minimumCost(cost []int) int {
    sort.Slice(cost, func(x, y int) bool {
        return cost[x] > cost[y]
    })

    var res int
    for k, v := range cost {
        if (k+1) % 3 != 0 {
            res += v
        }
    }

    return res
}
