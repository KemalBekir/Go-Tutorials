package main

import "strings"

func repeatedString(s string, n int64) int64 {
	var count int64
	count += int64(strings.Count(s, "a")) * (n / int64(len(s)))
	count += int64(strings.Count(string(s[:n%int64(len(s))]), "a"))
	return count
}
