package main

import "math"

func formingMagicSquare(s [][]int32) int32 {
	ms := [][]int32{
		{8, 1, 6, 3, 5, 7, 4, 9, 2}, {6, 1, 8, 7, 5, 3, 2, 9, 4},
		{4, 3, 8, 9, 5, 1, 2, 7, 6}, {2, 7, 6, 9, 5, 1, 4, 3, 8},
		{2, 9, 4, 7, 5, 3, 6, 1, 8}, {4, 9, 2, 3, 5, 7, 8, 1, 6},
		{6, 7, 2, 1, 5, 9, 8, 3, 4}, {8, 3, 4, 1, 5, 9, 6, 7, 2},
	}

	arr := flattenArray(&s)
	minCost := int32(math.MaxInt32)

	for _, m := range ms {
		cost := int32(0)

		for i := 0; i < len(m); i++ {
			cost += absolute(m[i] - arr[i])
		}

		minCost = compareMin(&minCost, &cost)
	}

	return minCost
}

func compareMin(a, b *int32) int32 {
	if *a < *b {
		return *a
	}

	return *b
}

func absolute(num int32) int32 {
	if num < 0 {
		return num * -1
	}

	return num
}

func flattenArray(s *[][]int32) []int32 {
	arr := []int32{}

	for i := 0; i < len(*s); i++ {
		for j := 0; j < len((*s)[0]); j++ {
			arr = append(arr, (*s)[i][j])
		}
	}

	return arr
}
