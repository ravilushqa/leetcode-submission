func longestSubarray(nums []int) int {
	var zeros, max, left int 

	for right, num := range nums {
        if num == 0 {
            zeros++
        }

        for zeros > 1 {
            if nums[left] == 0 {
                zeros--
            }
            
            left++
        }

        if right - left > max {
            max = right - left
        }
	}


	return max
}
