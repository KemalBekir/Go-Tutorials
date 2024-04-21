package main

func biggerIsGreater(w string) string {
	runes := []rune(w)

	i := len(runes) - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}

	if i < 0 {
		return "no answer"
	}

	j := len(runes) - 1
	for runes[j] <= runes[i] {
		j--
	}

	runes[i], runes[j] = runes[j], runes[i]

	reverse(runes[i+1:])

	return string(runes)
}

func reverse(runes []rune) {
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}
