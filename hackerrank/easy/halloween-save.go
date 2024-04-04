package main

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	count := int32(0)

	for s >= p {
		s -= p
		count++

		if p-d > m {
			p -= d
		} else {
			p = m
		}
	}
	return count
}
