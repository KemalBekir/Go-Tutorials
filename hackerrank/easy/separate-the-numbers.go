package main

import (
	"strconv"
	"fmt"
)

func separateNumbers(s string) {
	k := 1
	ans := false
	var n int64

	for k < len(s) {
		// Convert the first k characters to an integer
		num, err := strconv.ParseInt(s[:k], 10, 64)
		if err != nil {
			fmt.Println("NO")
			return
		}

		// Create the temporary string by appending increasing numbers
		temp := strconv.FormatInt(num, 10)
		i := int64(1)
		for len(temp) < len(s) {
			nextNum := num + i
			temp += strconv.FormatInt(nextNum, 10)
			i++
		}

		// Check if the constructed string matches the input string
		if temp == s {
			ans = true
			n = num
			break
		}

		k++
	}

	if ans {
		fmt.Printf("YES %d\n", n)
	} else {
		fmt.Println("NO")
	}
}