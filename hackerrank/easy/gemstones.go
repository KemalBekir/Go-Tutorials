package main

func gemstones(arr []string) int32 {
	if len(arr) == 0 {
		return 0
	}

	res := make(map[rune]bool)
	for _, char := range arr[0] {
		res[char] = true
	}

	for _, mineral := range arr[1:] {
		currentSet := make(map[rune]bool)
		for _, char := range mineral {
			currentSet[char] = true
		}

		for char := range res {
			if !currentSet[char] {
				delete(res, char)
			}
		}
	}

	return int32(len(res))
}