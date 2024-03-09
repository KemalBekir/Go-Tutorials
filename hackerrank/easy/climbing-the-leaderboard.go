package main

import "sort"

func uniqueValuedSlice(input []int32) []int32 {
	u := make([]int32, 0, len(input))
	m := make(map[int32]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	ranked = uniqueValuedSlice(ranked)
	sort.Slice(ranked, func(i, j int) bool { return ranked[i] < ranked[j] })
	var placings []int32
	var index int
	for _, s := range player {
		for len(ranked) > index && s >= ranked[index] {
			index++
		}
		placings = append(placings, int32(len(ranked)+1-index))
	}
	return placings
}
