func selfDividingNumbers(left int, right int) []int {
    res := make([]int, 0, right-left)
	for i := left; i <= right; i++ {
		num := i
		notValid := false
		for num != 0 {
			digit := num % 10
			if digit == 0 {
				notValid = true
				break
			}
			if i%digit != 0 {
				notValid = true
				break
			}
			num = num/10
		}

		if !notValid {
			res = append(res, i)
		}

	}

	return res
}
