package main

func superReducedString(s string) string {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			s = s[:i-1] + s[i+1:]
			return superReducedString(s)
		}
	}
	if len(s) == 0 {
		return "Empty String"
	}
	return s
}
