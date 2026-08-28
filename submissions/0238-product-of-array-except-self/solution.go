func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    // prefix
    product := 1
    for i := 0; i < len(nums); i++ {
        res[i] = product
        product *= nums[i]
    }

    // suffix
    product = 1
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= product
        product *= nums[i]
    }
    return res
}
