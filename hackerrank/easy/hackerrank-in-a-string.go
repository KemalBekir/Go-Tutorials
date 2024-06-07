package main

func hackerrankInString(s string) string {
	target := "hackerrank"
	j := 0

	for i := 0; i < len(s) && j < len(target); i++ {
		if s[i] == target[j] {
			j++
		}
	}

	if j == len(target) {
		return "YES"
	}
	return "NO"
}
