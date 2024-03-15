package main

import "math"

func viralAdvertising(n int32) int32 {
	var totalLikeCount int32 = 2

	if n > 1 {
		var lastLikeCount = totalLikeCount
		for i := 1; i < int(n); i++ {
			var newShares int32 = lastLikeCount * 3
			lastLikeCount = int32(math.Floor(float64(newShares / 2)))
			totalLikeCount += lastLikeCount
		}
	}

	return totalLikeCount
}
