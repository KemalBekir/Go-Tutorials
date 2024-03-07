package main

func countingValleys(steps int32, path string) int32 {
	var count int32
	var seaLevel int

	for _, v := range path {
		if seaLevel == 0 && v == 'D' {
			count++
		}

		if v == 'U' {
			seaLevel++
		} else if v == 'D' {
			seaLevel--
		}
	}

	return count
}
