package main

import "math"

func maxSubarray(arr []int32) []int32 {
	if len(arr) == 0 {
		return []int32{0, 0}
	}

	bestSum := int32(0)
	currentSum := int32(0)
	subSeqSum := int32(0)
	nmax := int32(math.MinInt32) // Smallest possible int32 value

	for _, x := range arr {
		if x > nmax {
			nmax = x
		}
		currentSum = max(0, currentSum+x)
		bestSum = max(bestSum, currentSum)
		if x > 0 {
			subSeqSum += x
		}
	}

	if subSeqSum == 0 { // If all numbers are negative or zero
		bestSum = nmax
		subSeqSum = nmax
	}

	return []int32{bestSum, subSeqSum}
}
