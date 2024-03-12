package main

func angryProfessor(k int32, a []int32) string {
	onTime := int32(0)
	for _, s := range a {
		if s <= 0 {
			onTime++
		}
	}

	if onTime >= k {
		return "NO"
	}
	return "YES"
}
