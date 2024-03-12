package main

import (
	"math"
	"strconv"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func beautifulDays(i int32, j int32, k int32) int32 {
	count := int32(0)

	for l := i; l <= j; l++ {
		reversedNum, _ := strconv.Atoi(reverseString(strconv.Itoa(int(l))))
		result := math.Abs(float64(l - int32(reversedNum)))

		if result/float64(k) == float64(int(result/float64(k))) {
			count++
		}
	}

	return count
}
