package main

import "sort"

func toys(w []int32) int32 {
    // Sort the slice
    sort.Slice(w, func(i, j int) bool { return w[i] < w[j] })

    var numContainers int32 = 0
    i := 0
    for i < len(w) {
        numContainers++
        containerLimit := w[i] + 4
        for i < len(w) && w[i] <= containerLimit {
            i++
        }
    }

    return numContainers
}