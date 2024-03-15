package main

import "sort"

func getTotalX(a []int32, b []int32) int32 {
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	minValue := a[0]
	maxValue := b[len(b)-1]
	var count int32

	for i := minValue; i <= maxValue; i++ {
		res := true
		for _, v := range a {
			res = res && i%v == 0
		}
		for _, v := range b {
			res = res && v%i == 0
		}
		if res != false {
			count++
		}
	}
	return count
}
