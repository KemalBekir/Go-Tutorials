package main

import (
	"fmt"
	"sort"
)

type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func almostSorted(arr []int32) {
	sortedArr := make([]int32, len(arr))
	copy(sortedArr, arr)
	sort.Sort(Int32Slice(sortedArr))

	var ans []int
	for i := range arr {
		if arr[i] != sortedArr[i] {
			ans = append(ans, i)
		}
	}

	if len(ans) == 0 {
		fmt.Println("yes")
	} else if len(ans) == 2 {
		fmt.Println("yes")
		fmt.Printf("swap %d %d\n", ans[0]+1, ans[1]+1)
	} else if len(ans) > 2 {
		reversible := true
		for j := 1; j < len(ans); j++ {
			if arr[ans[j-1]] < arr[ans[j]] {
				fmt.Println("no")
				reversible = false
				break
			}
		}
		if reversible {
			fmt.Println("yes")
			fmt.Printf("reverse %d %d\n", ans[0]+1, ans[len(ans)-1]+1)
		}
	} else {
		fmt.Println("no")
	}
}
