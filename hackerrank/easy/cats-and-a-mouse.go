package main

func abs(a int32, b int32) int32 {
	if a-b < 0 {
		return (a - b) * -1
	}
	return a - b
}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {

	if abs(z, x) < abs(z, y) {
		return "Cat A"
	} else if abs(z, x) == abs(z, y) {
		return "Mouse C"
	}

	return "Cat B"
}
