package main

func alternatingCharacters(s string) int32 {
	if len(s) == 0 {
		return 0
	}

	prev := s[0]
	var cnt int32 = 0

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			cnt++
		} else {
			prev = s[i]
		}
	}

	return cnt
}