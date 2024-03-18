package main

func jumpingOnClouds(c []int32, k int32) int32 {
	energy := int32(100)
	for i := int(k); i != 0; i += int(k) {
		if i >= len(c) {
			i = i - len(c)
		}
		energy--
		if c[i] == 1 {
			energy -= 2
		}
		if i == 0 {
			break
		}
	}
	return energy
}
