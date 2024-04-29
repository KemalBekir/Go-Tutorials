package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(arr []int) {
	N := len(arr)
	for i := 1; i < N; i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = value
	}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Reading input from stdin
	input := scanner.Text()
	N, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Error reading number of elements:", err)
		return
	}

	scanner.Scan()
	input = scanner.Text()
	inputs := strings.Fields(input)
	arr := make([]int, N)

	for i, v := range inputs {
		arr[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error reading array elements:", err)
			return
		}
	}

	insertionSort(arr)
}
