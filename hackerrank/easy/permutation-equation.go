package main

func permutationEquation(p []int32) []int32 {
	result := make([]int32, len(p))
	d := make([]int, len(p)+1)

	for i, v := range p {
		d[v] = i
	}

	for i := range p {
		result[i] = int32(d[d[i+1]+1] + 1)
	}

	return result
}
