package main

func icecreamParlor(m int32, arr []int32) []int32 {
    // Map to store the indices of the costs
    costMap := make(map[int32]int32)

    for i, cost := range arr {
        complement := m - cost
        if index, found := costMap[complement]; found {
            // Found the pair, return 1-based indices
            return []int32{index + 1, int32(i) + 1}
        }
        // Store the index of the current cost
        costMap[cost] = int32(i)
    }

    // Return an empty slice if no pair is found
    return []int32{}
}