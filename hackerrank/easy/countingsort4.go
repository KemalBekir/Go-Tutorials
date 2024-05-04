package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countSort(arr [][]string) {
	// Create a slice of slices to hold groups of strings based on the index in the array.
	res := make([][]string, len(arr)/2)

	// Populate the result slice based on input.
	for i, l := range arr {
		index := 0
		if idx, err := strconv.Atoi(l[0]); err == nil {
			index = idx
		}
		if index < len(res) { // Ensure the index is within the new slice's range.
			if i < len(arr)/2 {
				res[index] = append(res[index], "-")
			} else {
				res[index] = append(res[index], l[1])
			}
		}
	}

	// Construct the final output string.
	var output []string
	for _, subarr := range res {
		if len(subarr) > 0 {
			output = append(output, strings.Join(subarr, " "))
		}
	}

	fmt.Println(strings.Join(output, " "))
}
