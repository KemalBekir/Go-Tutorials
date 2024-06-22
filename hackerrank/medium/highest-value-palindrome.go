package main

import "strings"

func highestValuePalindrome(s string, n int32, k int32) string {
	sb := []rune(s)
	diff := make(map[int]bool)
	
	// First pass: make the string a palindrome
	for i := int32(0); i < n/2; i++ {
		if k < 0 {
			break
		}
		if sb[i] != sb[n-1-i] {
			m := max(sb[i], sb[n-1-i])
			sb[i] = m
			sb[n-1-i] = m
			diff[int(i)] = true
			k--
		}
	}

	if k < 0 {
		return "-1"
	}

	// Second pass: maximize the palindrome
	if k > 0 {
		for i := int32(0); i < n/2; i++ {
			if k == 0 {
				break
			}
			if sb[i] != '9' {
				if _, exists := diff[int(i)]; exists {
					sb[i] = '9'
					sb[n-1-i] = '9'
					k--
				} else if k > 1 {
					sb[i] = '9'
					sb[n-1-i] = '9'
					k -= 2
				}
			}
		}
	}

	// If there's one change left and the length of the string is odd, set the middle element to '9'
	if k >= 1 && n%2 == 1 {
		sb[n/2] = '9'
	}

	return string(sb)
}

func max(a, b rune) rune {
	if a > b {
		return a
	}
	return b
}