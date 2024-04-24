package main

import "sort"

func angryChildren(k int32, packets []int32) int64 {
	sort.Slice(packets, func(i, j int) bool { return packets[i] < packets[j] })
	n := len(packets)
	s := int64(-packets[0])
	d := int64(0)
	k64 := int64(k)

	for i := int64(0); i < k64; i++ {
		s += int64(packets[i])
		d += (2*i - k64 + 1) * int64(packets[i])
	}
	ans := d

	for shift := int64(1); shift < int64(n)-k64+1; shift++ {
		d += (-2*s + (k64-1)*(int64(packets[shift-1])+int64(packets[int(shift)+int(k64)-1])))
		s += int64(packets[int(shift)+int(k64)-1]) - int64(packets[shift])
		ans = min64(ans, d)
	}
	return ans
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
