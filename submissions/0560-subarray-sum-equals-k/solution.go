func subarraySum(nums []int, k int) int {
    res := 0
    prefix := map[int]int{0:1}

    // map 


    sum := 0
    for _, v := range nums {
        sum += v
        if prefix[sum - k] > 0 {
            res += prefix[sum - k]
        }

        prefix[sum]++
    }

    return res
}
