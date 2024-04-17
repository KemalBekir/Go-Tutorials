package main

import "strings"

func arm(grid []string) []string {
	full := make([]string, len(grid))
	for i := range grid {
		full[i] = strings.Repeat("O", len(grid[i]))
	}
	return full
}

func clear(g []string, i, j int) bool {
	if g[i][j] == 'O' ||
		(i > 0 && g[i-1][j] == 'O') ||
		(i < len(g)-1 && g[i+1][j] == 'O') ||
		(j > 0 && g[i][j-1] == 'O') ||
		(j < len(g[i])-1 && g[i][j+1] == 'O') {
		return true
	}
	return false
}

func explode(grid []string) []string {
	newGrid := make([]string, len(grid))
	for i, row := range grid {
		runes := make([]rune, len(row))
		for j, _ := range row {
			if clear(grid, i, j) {
				runes[j] = '.'
			} else {
				runes[j] = 'O'
			}
		}
		newGrid[i] = string(runes)
	}
	return newGrid
}

func bomberMan(n int32, grid []string) []string {
	if n == 1 {
		return grid
	}

	if n%2 == 0 {
		return arm(grid)
	}

	firstExplosion := explode(grid)
	if n%4 == 3 {
		return firstExplosion
	}
	return explode(firstExplosion)
}
