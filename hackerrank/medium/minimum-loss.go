package main

import (
	"sort"
	"math"
)

func minimumLoss(price []int64) int32 {
    n := len(price)
    
    // Create a slice of structs containing price and original index
    type priceIndex struct {
        price int64
        index int
    }
    priceWithIndices := make([]priceIndex, n)
    for i, p := range price {
        priceWithIndices[i] = priceIndex{p, i}
    }
    
    // Sort the slice by price
    sort.Slice(priceWithIndices, func(i, j int) bool {
        return priceWithIndices[i].price < priceWithIndices[j].price
    })
    
    minLoss := int64(math.MaxInt64)
    
    // Iterate through the sorted prices and calculate the minimum loss
    for i := 1; i < n; i++ {
        currentPrice := priceWithIndices[i].price
        currentIndex := priceWithIndices[i].index
        previousPrice := priceWithIndices[i-1].price
        previousIndex := priceWithIndices[i-1].index
        
        // Ensure that we buy before we sell
        if currentIndex < previousIndex {
            loss := currentPrice - previousPrice
            if loss < minLoss {
                minLoss = loss
            }
        }
    }
    
    return int32(minLoss)
}