package main

import "math"

func squares(a int32, b int32) int32 {
	left, right := int32(math.Sqrt(float64(a))), int32(math.Sqrt(float64(b)))
	if left*left < a {
		left++
	}
	return right - left + 1

}
