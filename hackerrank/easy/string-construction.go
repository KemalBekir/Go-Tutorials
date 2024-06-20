package main

func stringConstruction(s string) int32 {
    // Create a set (map with bool) for unique characters in the string
    charSet := make(map[rune]bool)
    for _, char := range s {
        charSet[char] = true
    }

    // The number of unique characters in the set is the cost
    return int32(len(charSet))
}
