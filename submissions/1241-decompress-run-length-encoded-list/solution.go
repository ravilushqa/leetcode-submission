func decompressRLElist(nums []int) []int {
	res := make([]int, 0, 1000)
	for i := range nums {
		if i >= len(nums)/2 {
			return res
		}
		freq, val := nums[2*i], nums[2*i+1]

		for freq != 0 {
			res=append(res,val)
			freq--
		}
	}

	return res
}
