package main

func pickingNumbers(a []int32) int32 {
	var maxCount int32 = 0

	for i := 0; i < len(a); i++ {
		count := int32(0)
		for j := 0; j < len(a); j++ {
			if a[i]-a[j] <= 1 && a[i]-a[j] >= 0 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
