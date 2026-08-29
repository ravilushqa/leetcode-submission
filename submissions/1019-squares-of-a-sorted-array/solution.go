func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    start, end := 0, len(nums) - 1

    for start <= end {
        startSqr := nums[start] * nums[start]
        endSqr := nums[end] * nums[end]
        if startSqr > endSqr {
            res[end-start] = startSqr
            start++
        } else {
            res[end-start] = endSqr
            end--
        }
    }

    return res
}
