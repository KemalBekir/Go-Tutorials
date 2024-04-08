package main

import "sort"

func flatlandSpaceStations(n int32, c []int32) int32 {
	if len(c) == int(n) {
		return 0
	}

	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	maximum := max(c[0], n-1-c[len(c)-1])

	for i := 1; i < len(c); i++ {
		distance := (c[i] - c[i-1]) / 2
		maximum = max(maximum, distance)
	}

	return maximum
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
