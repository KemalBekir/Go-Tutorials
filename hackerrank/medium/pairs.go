package main

func pairs(k int32, arr []int32) int32 {
    // Create a map for O(1) lookups
    arrMap := make(map[int32]bool)
    for _, v := range arr {
        arrMap[v] = true
    }
    
    count := int32(0)
    
    // Check for each element if there's a pair with difference k
    for _, x := range arr {
        if _, exists := arrMap[x+k]; exists {
            count++
        }
        if _, exists := arrMap[x-k]; exists {
            count++
        }
    }
    
    // Since each pair is counted twice, we return half the count
    return count / 2
}