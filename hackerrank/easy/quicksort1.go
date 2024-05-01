package main

func quickSort(arr []int32) []int32 {
	// Base case: if the array is of length 0 or 1, it is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Choose the first element as the pivot
	pivot := arr[0]
	var left []int32
	var right []int32

	// Partition the array into elements less than the pivot and elements greater than the pivot
	for _, value := range arr[1:] { // start from the second element
		if value < pivot {
			left = append(left, value)
		} else if value > pivot {
			right = append(right, value)
		}
	}

	// Recursively sort the left and right partitions and concatenate them with the pivot
	left = quickSort(left)
	right = quickSort(right)

	// Concatenate the sorted left slice, pivot, and sorted right slice
	sorted := append(left, pivot)
	sorted = append(sorted, right...)

	return sorted
}
