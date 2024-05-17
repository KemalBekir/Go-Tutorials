package main

func theLoveLetterMystery(s string) int32 {
	var total int32
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		// Compute the absolute difference between symmetric characters
		total += int32(absLetter(int(runes[i]) - int(runes[n-i-1])))
	}

	return total
}

// abs returns the absolute value of an integer
func absLetter(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
