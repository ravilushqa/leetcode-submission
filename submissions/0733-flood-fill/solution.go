func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    var bfs func(c,r int)

    directions := [][]int{{1,0},{-1,0},{0,1},{0,-1}}

    rows, cols := len(image), len(image[0])

    originColor := image[sr][sc]
    if originColor == color {
        return image
    }
    bfs = func(r,c int) {
        queue := []int{r,c}
        image[r][c] = color

        for len(queue) > 0 {
            r, c  := queue[0], queue[1]
            queue = queue[2:]

            for _, dir := range directions {
                row, col := r+dir[0], c+dir[1]

                if col >= 0 && row >= 0 && 
                col < cols && row < rows &&
                image[row][col] == originColor {
                    queue = append(queue, row, col)
                    image[row][col] = color
                }
            }
        }
    }

    bfs(sr, sc)

    return image
}
