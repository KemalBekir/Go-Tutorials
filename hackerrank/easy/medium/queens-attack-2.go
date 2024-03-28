package main

import "math"

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func sign(a int32) int32 {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	possible := [3][3]int32{}
	possible[0][1] = n - r_q                             // up
	possible[1][2] = n - c_q                             // right
	possible[1][0] = c_q - 1                             // left
	possible[2][1] = r_q - 1                             // down
	possible[0][0] = min(possible[0][1], possible[1][0]) // up left
	possible[0][2] = min(possible[0][1], possible[1][2]) // up right
	possible[2][0] = min(possible[2][1], possible[1][0]) // down left
	possible[2][2] = min(possible[2][1], possible[1][2]) // down right

	for i := int32(0); i < k; i++ {
		diffr := obstacles[i][0] - r_q
		diffc := c_q - obstacles[i][1]
		if diffr == 0 || diffc == 0 || absQueen(diffr) == absQueen(diffc) {
			possible[1-sign(diffr)][1-sign(diffc)] = min(possible[1-sign(diffr)][1-sign(diffc)], max(absQueen(diffr), absQueen(diffc))-1)
		}
	}

	sum := int32(0)
	for i := int32(0); i < 3; i++ {
		for j := int32(0); j < 3; j++ {
			sum += possible[i][j]
		}
	}
	return sum
}

func absQueen(a int32) int32 {
	return int32(math.Abs(float64(int(a))))
}
