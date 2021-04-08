func shuffle(nums []int, n int) []int {
    res := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		res = append(res, []int{nums[i], nums[i+n]}...)
	}
	return res
}
