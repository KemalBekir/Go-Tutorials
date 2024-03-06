package main

func migratoryBirds(arr []int32) int32 {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[int(v)]++
	}

	var maxCount int
	var birdType int

	for k, v := range counts {
		if maxCount < v || (maxCount == v && k < birdType) {
			maxCount = v
			birdType = k
		}
	}
	return int32(birdType)

}
