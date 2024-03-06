package main

func sockMerchant(n int32, ar []int32) int32 {
	counts := make(map[int]int)
	for _, v := range ar {
		counts[int(v)]++
	}

	totalPairs := 0
	for _, v := range counts {
		totalPairs += v / 2
	}

	return int32(totalPairs)
}
