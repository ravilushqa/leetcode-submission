func setZeroes(matrix [][]int)  {
    zeroRows := make(map[int]struct{}, len(matrix))
	zeroCols := make(map[int]struct{}, len(matrix[0]))
	for row := range matrix {
		for col, v := range matrix[row] {
			if v == 0 {
				zeroRows[row] = struct{}{}
				zeroCols[col] = struct{}{}
			}
		}
	}

	for zeroRow := range zeroRows {
		for i := range matrix[zeroRow] {
			matrix[zeroRow][i] = 0
		}
	}

	for row := range matrix {
		for zeroCol := range zeroCols {
			matrix[row][zeroCol] = 0
		}
	}
}
