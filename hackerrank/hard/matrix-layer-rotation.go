package main

import "fmt"

func matrixRotation(matrix [][]int32, r int32) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	k := min(m, n) / 2

	for layer := 0; layer < k; layer++ {
		var temp []int32

		// Top row
		for j := layer; j < n-layer-1; j++ {
			temp = append(temp, matrix[layer][j])
		}
		// Right column
		for i := layer; i < m-layer-1; i++ {
			temp = append(temp, matrix[i][n-layer-1])
		}
		// Bottom row
		for j := n - layer - 1; j > layer; j-- {
			temp = append(temp, matrix[m-layer-1][j])
		}
		// Left column
		for i := m - layer - 1; i > layer; i-- {
			temp = append(temp, matrix[i][layer])
		}

		// Rotate the temp array
		rotated := rotateArray(temp, int(r))

		idx := 0
		// Top row
		for j := layer; j < n-layer-1; j++ {
			matrix[layer][j] = rotated[idx]
			idx++
		}
		// Right column
		for i := layer; i < m-layer-1; i++ {
			matrix[i][n-layer-1] = rotated[idx]
			idx++
		}
		// Bottom row
		for j := n - layer - 1; j > layer; j-- {
			matrix[m-layer-1][j] = rotated[idx]
			idx++
		}
		// Left column
		for i := m - layer - 1; i > layer; i-- {
			matrix[i][layer] = rotated[idx]
			idx++
		}
	}

	// Print the matrix
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func rotateArray(arr []int32, r int) []int32 {
	l := len(arr)
	r %= l
	return append(arr[r:], arr[:r]...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
