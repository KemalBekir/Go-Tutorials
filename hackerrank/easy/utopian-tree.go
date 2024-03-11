package main

func utopianTree(n int32) int32 {
	var height int32 = 1

	for i := 0; i < int(n); i++ {
		if i%2 == 0 {
			height *= 2
		} else {
			height++
		}
	}
	return height
}
