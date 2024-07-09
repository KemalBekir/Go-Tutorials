package main

func shortPalindrome(s string) int32 {
    const mod = int64(1e9 + 7)
    arr1 := make([]int64, 26)
    arr2 := make([][]int64, 26)
    for i := range arr2 {
        arr2[i] = make([]int64, 26)
    }
    arr3 := make([]int64, 26)
    var ans int64

    for i := 0; i < len(s); i++ {
        idx := s[i] - 'a'
        ans = (ans + arr3[idx]) % mod
        for j := 0; j < 26; j++ {
            arr3[j] = (arr3[j] + arr2[j][idx]) % mod
        }
        for j := 0; j < 26; j++ {
            arr2[j][idx] = (arr2[j][idx] + arr1[j]) % mod
        }
        arr1[idx] = (arr1[idx] + 1) % mod
    }

    return int32(ans)
}