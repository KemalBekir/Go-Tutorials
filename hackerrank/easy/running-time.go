package main

func runningTime(arr []int32) int32 {
	var res int32 = 0
	for i := 1; i < len(arr); i++ {
		j := i - 1
		key := arr[i]

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
			res++
		}
		arr[j+1] = key
	}
	return res
}
