package main

func countingSort2(arr []int32) []int32 {
	if len(arr) == 0 {
		return nil
	}

	// Find the maximum value in the array to determine the size of the counting array
	var maxVal int32
	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
	}

	// Create a counting array with size (maxVal + 1)
	count := make([]int32, maxVal+1)

	// Fill the counting array with the frequency of each number
	for _, num := range arr {
		count[num]++
	}

	// Build the sorted array using the frequency count
	var result []int32
	for i, c := range count {
		for j := int32(0); j < c; j++ {
			result = append(result, int32(i))
		}
	}

	return result
}
