func orangesRotting(grid [][]int) int {
    directions := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
    queue := [][2]int{}
    freshOranges := 0
    timeElapsed := -1

    rows , cols := len(grid), len(grid[0])
    for r := range rows {
        for c := range cols {
            if grid[r][c] == 1 {
                freshOranges += 1
            } else if grid[r][c] == 2 {
                queue = append(queue, [2]int{r,c})
            }
        }
    }

    if freshOranges == 0 {
        return 0
    }
    

    for len(queue) > 0 {
        for _ = range len(queue) {
            el := queue[0]
            queue = queue[1:]
            
            r, c := el[0], el[1]

            for _, dir := range directions {
                row, col := dir[0] + r, dir[1] + c

                if row < 0 || col < 0 || row >= rows || col >= cols ||
                grid[row][col] != 1 {
                    continue
                }

                grid[row][col] = 2
                freshOranges--
                queue = append(queue, [2]int{row,col})
            }
        }
        timeElapsed++
    }

    if freshOranges > 0 {
        return -1
    } else {
        return timeElapsed
    }
}
