package main	

// Solution to Jumping on Clouds
func jumpClouds(c []int32) int32 {
	result := int32(0)

	for i := 1; i < len(c); i++ {
		if (i+1) < len(c) && c[i+1] != 1 {
			result++
			i++
			continue
		}
		if c[i] != 1 {
			result++
		}
	}
	return result
}