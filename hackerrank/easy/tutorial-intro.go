package main

func introTutorial(V int32, arr []int32) int32 {
	result := int32(0)

	for i := 0; i < len(arr); i++ {
		if arr[i] == V {
			result = int32(i)
			break
		}
	}
	return result
}
