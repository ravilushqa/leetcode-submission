func minSubArrayLen(target int, nums []int) int {
    res := 0
    start := 0
    
    sum := 0
    for end, v := range nums {
        sum += v

        for sum >= target {
            if res == 0 {
                res = end - start + 1
            } else {
                res = min(res, end - start + 1)
            }

            sum -= nums[start]
            start++
        } 
    }

    return res
}
