package main

import (
	"fmt"
	"strconv"
)

func kaprekarNumbers(p int32, q int32) {
	var result []int
	for i := p; i <= q; i++ {
		square := int64(i) * int64(i)
		squareString := strconv.FormatInt(square, 10)
		halfLength := len(squareString) / 2

		var num1, num2 int64
		var err error
		if halfLength > 0 {
			num1, err = strconv.ParseInt(squareString[:halfLength], 10, 64)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
		}
		num2, err = strconv.ParseInt(squareString[halfLength:], 10, 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}

		if num1+num2 == int64(i) {
			result = append(result, int(i))
		}
	}

	if len(result) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		for _, num := range result {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
