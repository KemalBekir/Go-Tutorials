package main

import "sort"

func maximumToys(prices []int32, k int32) int32 {
    // Sort the prices slice
    sort.Slice(prices, func(i, j int) bool {
        return prices[i] < prices[j]
    })

    var count int32 = 0
    for _, price := range prices {
        if k-price >= 0 {
            k -= price
            count++
        } else {
            break
        }
    }

    return count
}
