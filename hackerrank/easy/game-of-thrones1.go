package main

func gameOfThrones(s string) string {
    // Create a map to store the count of each character
    charCount := make(map[rune]int)

    // Count the occurrences of each character
    for _, char := range s {
        charCount[char]++
    }

    // Check if at most one character has an odd count
    oddCount := 0
    for _, count := range charCount {
        if count%2 != 0 {
            oddCount++
        }
    }

    // If more than one character has an odd count, it's not possible to rearrange
    if oddCount > 1 {
        return "NO"
    }

    return "YES"
}