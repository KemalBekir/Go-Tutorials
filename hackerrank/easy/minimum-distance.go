package main

import "math"

func minimumDistances(a []int32) int32 {
	result := []int32{}
	for i := 0; i < len(a); i++ {
		distance := int32(0)
		for j := i; j < len(a); j++ {
			if a[i] == a[j] && i != j {
				distance = int32(math.Abs(float64(i - j)))
				result = append(result, distance)
			}
		}
	}

	if len(result) != 0 {
		return min(result)
	} else {
		return -1
	}
}

func min(a []int32) int32 {
	minValue := a[0]

	for _, value := range a {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}
