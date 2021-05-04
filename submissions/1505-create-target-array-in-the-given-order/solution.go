func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(index))

	for i, idx := range index {
		target = append(target[:idx+1], target[idx:len(index)-1]...)
		target[idx] = nums[i]
	}

	return target
}
