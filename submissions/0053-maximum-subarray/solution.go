func maxSubArray(nums []int) int {
	var i, curSubArraySum, best int
	for i < len(nums) {
		if i == 0 {
			curSubArraySum = nums[i]
            best = nums[i]
			i++
			continue
		}

        // keep
		if curSubArraySum >= 0 {
            curSubArraySum +=nums[i]
        } else { // discard
            curSubArraySum = nums[i]
        }

        if curSubArraySum > best {
            best = curSubArraySum
        }


        i++
	}

	return best
}
