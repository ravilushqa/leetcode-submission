func findNumbers(nums []int) int {
	var count int
	for _, num := range nums {
		if len(strconv.Itoa(num)) % 2 == 0 {
			count++
		}
	}
	return count
}
