package main

import "sort"

func jimOrders(orders [][]int32) []int32 {
    // Create a slice to store the serve times and their indices
    type Order struct {
        index     int32
        serveTime int32
    }
    serveTimeWithIndex := make([]Order, len(orders))

    // Calculate total preparation time for each order
    for i, order := range orders {
        serveTimeWithIndex[i] = Order{
            index:     int32(i + 1),
            serveTime: order[0] + order[1],
        }
    }

    // Sort the slice based on serve time
    sort.Slice(serveTimeWithIndex, func(i, j int) bool {
        if serveTimeWithIndex[i].serveTime == serveTimeWithIndex[j].serveTime {
            return serveTimeWithIndex[i].index < serveTimeWithIndex[j].index
        }
        return serveTimeWithIndex[i].serveTime < serveTimeWithIndex[j].serveTime
    })

    // Extract customer order indices
    result := make([]int32, len(orders))
    for i, order := range serveTimeWithIndex {
        result[i] = order.index
    }

    return result
}