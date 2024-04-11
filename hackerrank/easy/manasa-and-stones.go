package main

func stones(n int32, a int32, b int32) []int32 {
	if a == b {
		return []int32{(n - 1) * a}
	}

	var results []int32
	maxVal := maxStones(a, b) * (n - 1)
	minVal := minStones(a, b) * (n - 1)
	step := absValue(a - b)

	for value := minVal; value <= maxVal; value += step {
		results = append(results, value)
	}

	return results
}

func maxStones(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func minStones(x, y int32) int32 {
	if x < y {
		return x
	}

	return y
}

func absValue(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
