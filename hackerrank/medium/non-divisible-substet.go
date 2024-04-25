package main

func nonDivisibleSubset(k int32, s []int32) int32 {
	remainderCount := make(map[int32]int32)
	for _, num := range s {
		rem := num % k
		remainderCount[rem]++
	}

	total := int32(0)
	seen := make(map[int32]bool)

	for rem, count := range remainderCount {
		if rem == 0 || (k%2 == 0 && rem == k/2) {
			total += 1
		} else if !seen[rem] {
			opposite := k - rem
			if countOpp, exists := remainderCount[opposite]; exists {
				if countOpp > count {
					total += countOpp
				} else {
					total += count
				}
				seen[opposite] = true
			} else {
				total += count
			}
			seen[rem] = true
		}
	}

	return total
}
