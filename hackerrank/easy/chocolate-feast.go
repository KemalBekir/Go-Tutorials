package main

func chocolateFeast(n int32, c int32, m int32) int32 {
	wrappers := n/c
	result := wrappers
	
	for wrappers >= m {
		exchangeBars := wrappers / m
		result += exchangeBars
		wrappers = exchangeBars + (wrappers % m)
	}
	return result
 }