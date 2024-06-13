package main

func funnyString(s string) string {
    n := len(s)
    for i := 0; i < n/2; i++ {
        if abs(int(s[i])-int(s[i+1])) != abs(int(s[n-i-1])-int(s[n-i-2])) {
            return "Not Funny"
        }
    }
    return "Funny"
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}