package main

func equalizeArray(arr []int32) int32 {
	result := map[int32]int32{}
	for _, n := range arr {
		result[n]++
	}

	var deleted, max int32

	for _, c := range result {
		if max < c {
			max = c
		}
		deleted += c
	}
	return deleted - max
}
