package main

func anagram(s string) int32 {
    length := len(s)
    if length%2 != 0 {
        return -1
    }
    half := length / 2

    tally := make(map[rune]int)
    for _, c := range s[:half] {
        tally[c]++
    }
    for _, c := range s[half:] {
        tally[c]--
    }

    var total [2]int
    for _, count := range tally {
        if count >= 0 {
            total[0] += count
        } else {
            total[1] -= count
        }
    }

    if total[0] > total[1] {
        return int32(total[0])
    }
    return int32(total[1])
}
