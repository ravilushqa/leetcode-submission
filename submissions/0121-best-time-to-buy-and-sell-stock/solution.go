func maxProfit(prices []int) int {
    best, minVal := 0, math.MaxInt
    for _, v := range prices {
        minVal = min(minVal, v)
        best = max(best, v-minVal)
    }
    return best
}
