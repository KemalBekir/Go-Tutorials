package main 

func balancedSums(arr []int32) string {
    totalSum := int32(0)
    for _, num := range arr {
        totalSum += num
    }

    leftSum := int32(0)
    for _, num := range arr {
        if leftSum == (totalSum - leftSum - num) {
            return "YES"
        }
        leftSum += num
    }

    return "NO"
}