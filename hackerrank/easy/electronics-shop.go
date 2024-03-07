package main

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	max := int32(-1)

	for _, k := range keyboards {
		for _, d := range drives {
			if k+d <= b && k+d > max {
				max = k + d
			}
		}
	}

	return max
}
