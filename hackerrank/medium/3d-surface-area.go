package main

func surfaceArea(A [][]int32) int32 {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}

	H := len(A)
	W := len(A[0])
	var res int32 = 2 * int32(H) * int32(W)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i == 0 {
				res += A[i][j]
			} else {
				res += max(0, A[i][j]-A[i-1][j]) // Front face
			}

			if i == H-1 {
				res += A[i][j]
			} else {
				res += max(0, A[i][j]-A[i+1][j])
			}

			if j == 0 {
				res += A[i][j]
			} else {
				res += max(0, A[i][j]-A[i][j-1])
			}

			if j == W-1 {
				res += A[i][j]
			} else {
				res += max(0, A[i][j]-A[i][j+1])
			}
		}
	}

	return res
}
