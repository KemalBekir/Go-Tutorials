package main

func larrysArray(A []int32) string {
	for searchVal := int32(1); searchVal <= int32(len(A)); searchVal++ {
		currPos := indexOf(A, searchVal)
		for A[searchVal-1] != searchVal {
			if currPos-(searchVal-1) >= 2 {
				midPos := A[currPos-1]
				firstPos := A[currPos-2]
				A[currPos-2] = A[currPos]
				A[currPos-1] = firstPos
				A[currPos] = midPos
				currPos -= 2
			} else {
				if currPos+1 >= int32(len(A)) {
					return "NO"
				}
				lastPos := A[currPos+1]
				firstPos := A[currPos-1]
				A[currPos-1] = A[currPos]
				A[currPos+1] = firstPos
				A[currPos] = lastPos
				currPos -= 1
			}
		}
	}

	prev := int32(-1)
	for _, i := range A {
		if prev > i {
			return "NO"
		}
		prev = i
	}
	return "YES"
}

func indexOf(slice []int32, val int32) int32 {
	for i, v := range slice {
		if v == val {
			return int32(i)
		}
	}
	return -1
}
