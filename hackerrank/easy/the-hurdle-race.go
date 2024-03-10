package main

func hurdleRace(k int32, height []int32) int32 {
	potion := int32(0)
	max := height[0]

	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}

	if max > k {
		potion = max - k
		return potion
	}

	return potion
}
