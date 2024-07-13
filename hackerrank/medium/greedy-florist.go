package main

import "sort"

func getMinimumCost(k int32, c []int32) int32 {
    nFlowers := int32(len(c))
    
    // Sort c in descending order
    sort.Slice(c, func(i, j int) bool {
        return c[i] > c[j]
    })
    
    var price, count int32
    
    for i := int32(0); i < nFlowers; i++ {
        if i%k == 0 && i != 0 {
            count++
        }
        price += c[i] * (count + 1)
    }
    
    return price
}