package main

import "sort"

func maxMin(k int32, arr []int32) int32 {
    // Convert arr to []int for easier sorting
    intArr := make([]int, len(arr))
    for i, v := range arr {
        intArr[i] = int(v)
    }
    
    // Sort the array
    sort.Ints(intArr)
    
    // Initialize minUnfairness with the maximum possible value
    minUnfairness := int32(^uint32(0) >> 1)  // Max int32 value
    
    // Find the minimum unfairness
    for i := 0; i <= len(intArr)-int(k); i++ {
        unfairness := int32(intArr[i+int(k)-1] - intArr[i])
        if unfairness < minUnfairness {
            minUnfairness = unfairness
        }
    }
    
    return minUnfairness
}