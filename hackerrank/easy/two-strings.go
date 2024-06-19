package main

func twoStrings(s1 string, s2 string) string {
    // Create a set (map with bool) for characters in s1
    charSet := make(map[rune]bool)
    for _, char := range s1 {
        charSet[char] = true
    }

    // Check if any character in s2 is in the set
    for _, char := range s2 {
        if charSet[char] {
            return "YES"
        }
    }

    return "NO"
}