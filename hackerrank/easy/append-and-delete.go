package main

import "strings"

func appendAndDelete(s string, t string, k int32) string {
	if s == t {
		return "Yes"
	}
	for k > 0 {
		if strings.HasPrefix(s, t) {
			s = strings.TrimPrefix(s, t)
			if 2*len(t)+len(s) <= int(k) || (int(k) >= len(s) && (int(k)-len(s))%2 == 0) {
				return "Yes"
			}
			return "No"
		}
		t = string([]rune(t)[:len(t)-1])
		k--
	}
	return "No"
}
