package main

import "strings"

func happyLadybugs(b string) string {
	counts := make(map[rune]int)
	for _, char := range b {
		counts[char]++
	}
	for char, count := range counts {
		if char != '_' && count == 1 {
			return "NO"
		}
	}

	if !strings.ContainsRune(b, '_') {
		n := len(b)
		if n > 1 {
			for i := 1; i < n-1; i++ {
				if b[i] != b[i-1] && b[i] != b[i+1] {
					return "NO"
				}
			}
			if b[0] != b[1] || b[n-1] != b[n-2] {
				return "NO"
			}
		}
	}

	return "YES"
}
