package main

import "sort"

func findMedian(arr []int32) int32 {
	// Sort the array.
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Determine the median.
	length := len(arr)
	if length%2 != 0 {
		// Return the middle element for odd length arrays.
		return arr[length/2]
	} else {
		// Return the average of the two middle elements for even length arrays.
		// This calculation does integer division and could lead to truncation,
		// reflecting the requirements of the function signature.
		return (arr[length/2] + arr[length/2-1]) / 2
	}
}
