func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	var res string
	rows := make(map[int][]string, numRows)
	for i := range rows {
		rows[i] = make([]string, 0, numRows)
	}

	if numRows == 2 {
		for i, char := range s {
			rows[i%2] = append(rows[i%2], string(char))
		}

		for i := 0; i < numRows; i++ {
			res += strings.Join(rows[i], "")
		}

		return res
	}

	chunks := make([]string, 0, numRows)
	for i := 0; i < len(s)-1; i = i + numRows - 1 {
		lastIndex := i + numRows
		if lastIndex > len(s) {
			lastIndex = len(s)
		}
		chunks = append(chunks, s[i:lastIndex])
	}

	for i, chunk := range chunks {
		direction := i % 2 // 0 - forward; 1 - backward

		for j, char := range chunk {
			if direction == 0 {
				rows[j] = append(rows[j], string(char))
				continue
			}
			if direction == 1 {
				if j == 0 {
					continue
				}
				if j == numRows-1 && i < len(chunks)-1 {
					continue
				}
				rows[numRows-1-j] = append(rows[numRows-1-j], string(char))
				continue
			}
		}
	}

	for i := 0; i < numRows; i++ {
		res += strings.Join(rows[i], "")
	}

	return res
}

