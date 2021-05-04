func judgeCircle(moves string) bool {
	moveMap := make(map[rune]int, 4)
	for _, move := range moves {
		moveMap[move]++
	}

    return (moveMap[[]rune("U")[0]] == moveMap[[]rune("D")[0]]) && (moveMap[[]rune("L")[0]] == moveMap[[]rune("R")[0]])
}
