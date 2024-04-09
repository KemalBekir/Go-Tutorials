package main

import "strconv"

func fairRations(B []int32) string {
	b := make([]int32, len(B))
	sum := int32(0)
	for i, n := range B {
		b[i] = n % 2
		sum += b[i]
	}

	if sum%2 != 0 {
		return "NO"
	}

	var odds []int
	for i, val := range b {
		if val == 1 {
			odds = append(odds, i)
		}
	}

	var res int32 = 0
	i := 0

	for i < len(odds) {
		res += int32(odds[i+1]-odds[i]) * 2
		i += 2
	}

	return strconv.Itoa(int(res))
}
