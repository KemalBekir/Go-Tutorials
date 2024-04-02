package main

func beautifulTriplets(d int32, arr []int32) int32 {
	count := int32(0)
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i]+d == arr[j] {
				for k := j + 1; k < len(arr); k++ {
					if arr[j]+d == arr[k] {
						count++
						break
					}
				}
			}
		}
	}
	return count
}
