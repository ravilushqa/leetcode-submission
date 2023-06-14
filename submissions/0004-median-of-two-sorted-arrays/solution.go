func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)

	if len(merged) == 0 {
		return 0
	}

	if len(merged)%2 == 1 {
		return float64(merged[len(merged)/2])
	}

	return (float64(merged[len(merged)/2-1]) + float64(merged[len(merged)/2])) / 2
}
