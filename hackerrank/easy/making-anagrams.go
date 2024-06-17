package main

func makingAnagrams(s1 string, s2 string) int32 {
    // Create frequency maps for both strings
    freqMap1 := make(map[rune]int)
    freqMap2 := make(map[rune]int)

    // Count the frequency of each character in both strings
    for _, char := range s1 {
        freqMap1[char]++
    }
    for _, char := range s2 {
        freqMap2[char]++
    }

    var deletions int32

    // Iterate over the first frequency map
    for char, count := range freqMap1 {
        // If the character exists in the second map
        if count2, exists := freqMap2[char]; exists {
            // Subtract the minimum count from both maps
            deletions += int32(abs(count - count2))
            delete(freqMap2, char)
        } else {
            // If the character doesn't exist in the second map,
            // all occurrences need to be deleted from the first string
            deletions += int32(count)
        }
    }

    // Add the remaining counts from the second map
    for _, count := range freqMap2 {
        deletions += int32(count)
    }

    return deletions
}

// Helper function to get the absolute value of an integer
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
