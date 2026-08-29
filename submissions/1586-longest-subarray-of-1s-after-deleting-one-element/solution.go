func longestSubarray(nums []int) int {
    var start, res, zeros int

    for end, v := range nums {
        if v == 0 {
            zeros++
        }
        for zeros > 1 {
            if nums[start] == 0 {
                zeros--
            }
            start++
        }

        res = max(res, end - start + 1 - 1)
    }

    return res
}
