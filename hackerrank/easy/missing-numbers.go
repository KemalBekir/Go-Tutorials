package main

import "sort"

func missingNumbers(arr []int32, brr []int32) []int32 {
    // Create maps to count the frequency of each element in both arrays
    arrCounter := make(map[int32]int)
    brrCounter := make(map[int32]int)

    for _, num := range arr {
        arrCounter[num]++
    }

    for _, num := range brr {
        brrCounter[num]++
    }

    // Find the missing numbers
    missing := []int32{}
    for num, count := range brrCounter {
        if count > arrCounter[num] {
            missing = append(missing, num)
        }
    }

    // Sort the list of missing numbers
    sort.Slice(missing, func(i, j int) bool { return missing[i] < missing[j] })

    return missing
}