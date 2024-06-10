package main

import "strings"

func pangrams(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	str := strings.ToLower(s)

	for _, letter := range alphabet {
		if !strings.ContainsRune(str, letter) {
			return "not pangram"
		}
	}
	return "pangram"
}
