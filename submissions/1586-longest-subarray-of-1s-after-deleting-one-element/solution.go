func longestSubarray(nums []int) int {
    l := 0
    zeros := 0
    max := 0

    for r, v := range nums {
        if v == 0 {
            zeros++
        }

        for zeros > 1 {
            if nums[l] == 0 {
                zeros--
            }
            l++
        }

        if r - l + 1 - 1 > max {
            max = r - l + 1 - 1
        }
    }

    return max
}
