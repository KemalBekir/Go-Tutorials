package main

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	l := int32(len(a))
	result := make([]int32, len(queries))

	for index, val := range queries {
		pos := (val + l - k) % l
		if pos < 0 {
			pos += l
		}
		result[index] = a[pos]
	}
	return result
}
