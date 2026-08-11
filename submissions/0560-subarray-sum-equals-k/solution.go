func subarraySum(nums []int, k int) int {
    var res, cur int

    seen := map[int]int{0:1}

    // [1,1,1] k = 2
    // 0,1,2,3
    // 1 -2 0 | 0:1, 1:1
    // 2 0 +1 | 0:1,1:1,2:1
    // 3 1 +1 | 0:1,1:1,2:1,3:1

    for _, v := range nums {
        cur += v
        res += seen[cur-k]

        seen[cur]++
    }

    return res
}
