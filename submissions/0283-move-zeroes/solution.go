func moveZeroes(nums []int) []int {
    var left int
    for right, v := range nums {
        if v != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
    }

    return nums 
}
