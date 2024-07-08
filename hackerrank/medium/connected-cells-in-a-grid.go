package main

func connectedCell(matrix [][]int32) int32 {
    n := len(matrix)
    m := len(matrix[0])
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, m)
    }

    var dfs func(i, j int) int32
    dfs = func(i, j int) int32 {
        stack := [][2]int{{i, j}}
        var count int32 = 0
        for len(stack) > 0 {
            x, y := stack[len(stack)-1][0], stack[len(stack)-1][1]
            stack = stack[:len(stack)-1]

            if x < 0 || x >= n || y < 0 || y >= m || visited[x][y] || matrix[x][y] == 0 {
                continue
            }

            visited[x][y] = true
            count++

            directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
            for _, dir := range directions {
                stack = append(stack, [2]int{x + dir[0], y + dir[1]})
            }
        }
        return count
    }

    var maxRegion int32 = 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if matrix[i][j] == 1 && !visited[i][j] {
                sizeOfRegion := dfs(i, j)
                if sizeOfRegion > maxRegion {
                    maxRegion = sizeOfRegion
                }
            }
        }
    }

    return maxRegion
}