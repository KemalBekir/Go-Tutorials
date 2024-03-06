package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	sum := int32(0)
	total := int32(0)
	for i, num := range bill {
		total += num
		if i != int(k) {
			sum += num
		}
	}

	if (sum/2)-b != 0 {
		if (sum/2)-b < 0 {
			sum = ((sum / 2) - b) * -1
			fmt.Printf("%d", sum)
			return
		}
		fmt.Printf("%d", sum)
		return
	} else if ((sum / 2) - b) == 0 {
		fmt.Println("Bon Appetit")
		return
	}
}
