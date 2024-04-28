package main

import (
	"fmt"
	"sort"
)

func insertionSort2(n int32, arr []int32) {
	for i := 1; i < int(n); i++ {
		sort.Slice(arr[:i+1], func(j, k int) bool {
			return arr[j] < arr[k]
		})
		printArray2(arr)
	}
}

// printArray is a helper function to print the array
func printArray2(arr []int32) {
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
