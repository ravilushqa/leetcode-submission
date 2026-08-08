func longestSubarray(nums []int) int {
    var res, zeros, start int

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

        res = max(res, end-start)
    }

    return res
}

