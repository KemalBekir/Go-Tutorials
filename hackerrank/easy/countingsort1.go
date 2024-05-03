package main

func countingSort(arr []int32) []int32 {
	// Initialize the frequency array with size 100, set to 0
	frequency := make([]int32, 100)

	// Iterate over each number in the input array and increment the
	// corresponding index in the frequency array
	for _, num := range arr {
		frequency[num]++
	}

	// Return the frequency array
	return frequency
}
