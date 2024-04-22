package main

// In the main function in this problem
// we have to add width as an argument to our function
// "result := serviceLane(n,width,cases)"

func serviceLane(n int32, width []int32, cases [][]int32) []int32 {
	results := make([]int32, len(cases))

	for idx, c := range cases {
		start := c[0]
		end := c[1]
		minWidth := width[start]

		for j := start + 1; j <= end; j++ {
			if width[j] < minWidth {
				minWidth = width[j]
			}
		}

		results[idx] = minWidth
	}

	return results
}
