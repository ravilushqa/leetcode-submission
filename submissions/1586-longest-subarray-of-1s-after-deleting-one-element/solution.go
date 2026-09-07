func longestSubarray(nums []int) int {
    res := 0
    start := 0
    zeros := 0

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
