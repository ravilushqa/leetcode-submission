func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1

	for left < right {
		width := right - left
		min := minNum(height[left], height[right])

		if max < min*width {
			max = min * width
		}

		if height[right] > height[left] {
			left++
			continue
		}

		right--
	}

	return max
}

func minNum(a, b int) int {
	if a < b {
		return a
	}

	return b
}
