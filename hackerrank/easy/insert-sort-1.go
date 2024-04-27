package main

import "fmt"

func insertionSort1(n int32, arr []int32) {
	lastValue := arr[n-1]
	i := n - 2

	for i >= 0 && arr[i] > lastValue {
		arr[i+1] = arr[i]
		printArray(arr)
		i--
	}

	arr[i+1] = lastValue
	printArray(arr)
}

func printArray(arr []int32) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
