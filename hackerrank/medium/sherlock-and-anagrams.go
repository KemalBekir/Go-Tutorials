package main

import "sort"

func sherlockAndAnagrams(s string) int32 {
	count := int32(0)
	substringMap := make(map[string]int32)

	// Generate all substrings
	for start := 0; start < len(s); start++ {
		for end := start + 1; end <= len(s); end++ {
			substr := s[start:end]
			// Sort the substring to generate a signature
			signature := sortedString(substr)
			substringMap[signature]++
		}
	}

	// Count pairs of anagrammatic substrings
	for _, val := range substringMap {
		if val > 1 {
			count += val * (val - 1) / 2
		}
	}

	return count
}

// Helper function to sort characters in a string
func sortedString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}