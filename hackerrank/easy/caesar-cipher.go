package main

func caesarCipher(s string, k int32) string {
	var res string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += string((char-'a'+k)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			res += string((char-'A'+k)%26 + 'A')
		} else {
			res += string(char)
		}
	}
	return res
}
