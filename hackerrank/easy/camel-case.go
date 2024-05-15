package main

func camelcase(s string) int32 {
	// Initialize count to 1, as we are starting with the assumption that the first letter is lowercase
	var count int32 = 1

	// Loop through each character in the string
	for _, char := range s {
		// Check if the character is an uppercase letter
		if char >= 'A' && char <= 'Z' {
			count++
		}
	}

	return count
}
