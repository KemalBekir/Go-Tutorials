package main

import "sort"

func twoPluses(grid []string) int32 {
	type coordinate struct{ x, y int }
	var pluses [][]coordinate

	// Check if it's possible to form a plus centered at each 'G'
	for n := 1; n < len(grid)-1; n++ {
		for m := 1; m < len(grid[0])-1; m++ {
			if grid[n][m] == 'G' && grid[n][m-1] == 'G' && grid[n][m+1] == 'G' && grid[n-1][m] == 'G' && grid[n+1][m] == 'G' {
				plus := []coordinate{{n - 1, m}, {n, m - 1}, {n, m}, {n, m + 1}, {n + 1, m}}
				pluses = append(pluses, plus)
				i := 2
				for n+i < len(grid) && n-i >= 0 && m+i < len(grid[0]) && m-i >= 0 &&
					grid[n+i][m] == 'G' && grid[n-i][m] == 'G' && grid[n][m+i] == 'G' && grid[n][m-i] == 'G' {
					plus = append(plus, coordinate{n - i, m}, coordinate{n, m - i}, coordinate{n, m + i}, coordinate{n + i, m})
					pluses = append(pluses, append([]coordinate(nil), plus...))
					i++
				}
			}
		}
	}

	// Sort pluses by size, largest first
	sort.Slice(pluses, func(i, j int) bool {
		return len(pluses[i]) > len(pluses[j])
	})

	countG := 0
	for _, row := range grid {
		for _, c := range row {
			if c == 'G' {
				countG++
			}
		}
	}

	if len(pluses) == 0 && countG >= 2 {
		return 1
	} else if len(pluses) == 1 {
		if len(pluses[0]) == countG {
			return 1
		} else {
			return int32(len(pluses[0]))
		}
	} else {
		ans := len(pluses[0])
		for j, x := range pluses {
			for _, y := range pluses[j+1:] {
				allClear := true
				for _, h := range x {
					for _, z := range y {
						if h == z {
							allClear = false
							break
						}
					}
					if !allClear {
						break
					}
				}
				if allClear && len(x)*len(y) > ans {
					ans = len(x) * len(y)
				}
			}
		}
		return int32(ans)
	}
}
