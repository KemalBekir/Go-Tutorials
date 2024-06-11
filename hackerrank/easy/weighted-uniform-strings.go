package main

func weightedUniformStrings(s string, queries []int32) []string {
	// Create the alphabet map with weights
	alphabet := "_abcdefghijklmnopqrstuvwxyz"
	weights := make(map[rune]int)
	for i, char := range alphabet {
		weights[char] = i
	}

	// Calculate weights of uniform substrings
	weightSet := make(map[int]struct{})
	var previous rune
	var length int
	for _, letter := range s {
		if letter == previous {
			length++
		} else {
			previous = letter
			length = 1
		}
		weightSet[weights[letter]*length] = struct{}{}
	}

	// Evaluate the queries
	results := make([]string, len(queries))
	for i, query := range queries {
		if _, found := weightSet[int(query)]; found {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}
	return results
}