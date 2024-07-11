package main

func largestPermutation(k int32, arr []int32) []int32 {
    n := int32(len(arr))
    positionMap := make(map[int32]int32)
    
    // Initialize positionMap and find the maximum value
    currentMax := int32(0)
    for i, value := range arr {
        positionMap[value] = int32(i)
        if value > currentMax {
            currentMax = value
        }
    }
    
    for i := int32(0); i < n; i++ {
        if k == 0 {
            break
        }
        
        if arr[i] != currentMax {
            // Swap arr[i] and currentMax
            arr[positionMap[currentMax]], arr[i] = arr[i], currentMax
            
            // Update positionMap
            positionMap[arr[positionMap[currentMax]]] = positionMap[currentMax]
            positionMap[currentMax] = i
            
            // Decrement k
            k--
        }
        
        // Update currentMax to the next largest number
        currentMax--
    }
    
    return arr
}