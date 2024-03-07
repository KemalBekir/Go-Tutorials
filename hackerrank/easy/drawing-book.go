package main

func pageCount(n int32, p int32) int32 {
	front := p / 2
	back := (n - p) / 2
	if n%2 == 0 && p%2 == 1 && n > p {
		back++
	}

	if front < back {
		return front
	}
	return back

}
