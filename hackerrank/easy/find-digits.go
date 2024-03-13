package main

import "strconv"

func findDigits(n int32) int32 {
	runes := []rune(strconv.Itoa(int(n)))
	count := int32(0)
	for i := 0; i < len(runes); i++ {
		str := string(runes[i])
		num, _ := strconv.Atoi(str)
		if num != 0 {
			if n%int32(num) == 0 {
				count++
			}
		}
	}

	return count
}
