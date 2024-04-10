package main

func cavityMap(grid []string) []string {
	n := len(grid)
	result := make([]string, n)

	for i := 0; i < n; i++ {
		temp := ""
		for j := 0; j < n; j++ {
			cur := grid[i][j]

			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				temp += string(cur)
			} else if grid[i+1][j] >= cur || grid[i-1][j] >= cur || grid[i][j+1] >= cur || grid[i][j-1] >= cur {
				temp += string(cur)

			} else {
				temp += "X"
			}
		}
		result[i] = temp
	}

	return result
}
