package main

import "regexp"

func minimumNumber(n int32, password string) int32 {
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasChar := regexp.MustCompile(`[\!@#\$\%\^&*\(\)\-\+]`).MatchString(password)

	var toAdd int32
	if !hasLower {
		toAdd++
	}
	if !hasUpper {
		toAdd++
	}
	if !hasDigit {
		toAdd++
	}
	if !hasChar {
		toAdd++
	}

	if n-6 < 0 && 6-n > toAdd {
		toAdd = 6 - n
	}

	return toAdd
}
