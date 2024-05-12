package main

import "sort"

func closestNumbers(arr []int32) []int32 {
	// Sort the array.
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Initialize the minimum difference to the largest possible value.
	minDiff := int32(2147483647) // This represents the largest int32 positive value
	var result []int32

	// Iterate through the sorted array to find the minimum difference.
	for i := 0; i < len(arr)-1; i++ {
		absolute := absNum(arr[i] - arr[i+1])
		if absolute < minDiff {
			minDiff = absolute
			// Reset result with the new closest pair.
			result = []int32{arr[i], arr[i+1]}
		} else if absolute == minDiff {
			// Extend result with the current closest pair.
			result = append(result, arr[i], arr[i+1])
		}
	}

	return result
}

// Helper function to calculate absolute value for int32.
func absNum(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}
