func longestSubarray(nums []int) int {
    res := 0
    start := 0
    zerosCount := 0

    for end, v := range nums {
        if v == 0 {
            zerosCount++
        }

        for zerosCount > 1 {
            if nums[start] == 0 {
                zerosCount--
            }
            start++
        }

        res = max(res, end - start + 1 - 1)
    }

    return res
}
