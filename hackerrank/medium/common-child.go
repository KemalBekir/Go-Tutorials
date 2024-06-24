package main

func commonChild(s1 string, s2 string) int32 {
	n := len(s1)
	count := make([][]int32, n+1)
	for i := range count {
		count[i] = make([]int32, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				count[i][j] = count[i-1][j-1] + 1
			} else {
				count[i][j] = max(count[i-1][j], count[i][j-1])
			}
		}
	}

	return count[n][n]
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}