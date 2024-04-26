package main

import "sort"

func bigSorting(unsorted []string) []string {
	lengthMap := make(map[int][]string)
	var lengths []int

	for _, num := range unsorted {
		l := len(num)
		if _, ok := lengthMap[l]; !ok {
			lengths = append(lengths, l)
		}
		lengthMap[l] = append(lengthMap[l], num)
	}

	sort.Ints(lengths)

	var result []string
	for _, l := range lengths {
		sort.Slice(lengthMap[l], func(i, j int) bool {
			return lengthMap[l][i] < lengthMap[l][j]
		})
		result = append(result, lengthMap[l]...)
	}

	return result
}
