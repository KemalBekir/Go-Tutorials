package main

func workbook(n int32, k int32, arr []int32) int32 {
	pageNumber := int32(0)
	amountOfSpecialProblems := int32(0)

	for i := int32(0); i < n; i++ {
		problemsInChapter := arr[i]
		pageNumber++

		for j := int32(1); j <= problemsInChapter; j++ {
			if j == pageNumber {
				amountOfSpecialProblems++
			}

			if j%k == 0 && j != problemsInChapter {
				pageNumber++
			}
		}
	}

	return amountOfSpecialProblems
}

//Solved
