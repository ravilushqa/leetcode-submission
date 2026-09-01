func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    l, r := 0 , len(nums) - 1

    for l <= r {
        lsqr := nums[l] * nums[l]
        rsqr := nums[r] * nums[r]

        if lsqr > rsqr {
            res[r-l] = lsqr
            l++
        } else {
            res[r-l] = rsqr
            r--
        }
    }

    return res
}
