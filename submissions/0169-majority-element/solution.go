func majorityElement(nums []int) int {
    slices.Sort(nums)
    m := (len(nums) - 1) / 2
    return nums[m]
}
