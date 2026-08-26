func sortedSquares(nums []int) []int {
    for i, v := range nums {
        nums[i] *= v
    }

    sort.Slice(nums, func(i,j int) bool {
        return nums[i] < nums[j]
    })

    return nums
}
