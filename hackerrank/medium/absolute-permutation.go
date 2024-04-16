package main

func absolutePermutation(n int32, k int32) []int32 {
	if k == 0 {
		result := make([]int32, n)
		for i := range result {
			result[i] = int32(i + 1)
		}
		return result
	}

	if n%(2*k) != 0 {
		return []int32{-1}
	}

	permutation := make([]int32, 0, n)
	for i := int32(1); i <= n; i += 2 * k {
		for j := i + k; j < i+2*k; j++ {
			permutation = append(permutation, j)
		}
		for j := i; j < i+k; j++ {
			permutation = append(permutation, j)
		}
	}

	return permutation
}
