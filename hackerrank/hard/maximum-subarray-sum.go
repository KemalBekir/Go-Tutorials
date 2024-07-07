package main

func maximumSum(a []int64, m int64) int64 {
    prefixSum := int64(0)
    maxModuloSum := int64(0)
    prefixSums := make([]int64, 0)

    for _, num := range a {
        // Update the prefix sum modulo m
        prefixSum = (prefixSum + num) % m

        // Update maxModuloSum with the current prefixSum
        if prefixSum > maxModuloSum {
            maxModuloSum = prefixSum
        }

        // Binary search to find the smallest prefix sum greater than the current prefixSum
        lo := 0
        hi := len(prefixSums)
        for lo < hi {
            mid := (lo + hi) / 2
            if prefixSums[mid] > prefixSum {
                hi = mid
            } else {
                lo = mid + 1
            }
        }

        if lo < len(prefixSums) {
            // There exists a prefix sum greater than the current one
            potentialMax := (prefixSum - prefixSums[lo] + m) % m
            if potentialMax > maxModuloSum {
                maxModuloSum = potentialMax
            }
        }

        // Insert the current prefixSum into the sorted list
        prefixSums = append(prefixSums, 0)
        copy(prefixSums[lo+1:], prefixSums[lo:])
        prefixSums[lo] = prefixSum
    }

    return maxModuloSum
}