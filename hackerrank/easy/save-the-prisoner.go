package main

func saveThePrisoner(n int32, m int32, s int32) int32 {
	start := s
	max := n
	candy := m
	seat := start + (candy-1)%max
	if seat > max {
		seat -= max
	}
	return seat
}
