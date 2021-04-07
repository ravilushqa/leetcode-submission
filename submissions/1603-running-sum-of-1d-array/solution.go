func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	for i:=len(nums)-1; i>=0;i-- {
		for j := i; j>=0;j-- {
			res[i]+=nums[j]
		}
	}

	return res
}
