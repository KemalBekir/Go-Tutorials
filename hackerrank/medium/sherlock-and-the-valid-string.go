package main

import "sort"

func isValid(s string) string {
    // Create a frequency map of characters in the string
    charFreq := make(map[rune]int)
    for _, char := range s {
        charFreq[char]++
    }

    // Create a frequency map of the frequencies
    freqFreq := make(map[int]int)
    for _, freq := range charFreq {
        freqFreq[freq]++
    }

    // Convert the freqFreq map to a slice of pairs
    type pair struct {
        freq  int
        count int
    }
    var freqList []pair
    for k, v := range freqFreq {
        freqList = append(freqList, pair{k, v})
    }

    // Sort the list by the frequency count
    sort.Slice(freqList, func(i, j int) bool {
        return freqList[i].count > freqList[j].count
    })

    // Get the most common frequency
    target := freqList[0].freq

    // Remove the target frequency from the list
    freqList = freqList[1:]

    // If there are no other frequencies left, return "YES"
    if len(freqList) == 0 {
        return "YES"
    }

    // If there is exactly one other frequency, check the conditions
    if len(freqList) == 1 {
        other := freqList[0]
        if (abs(target-other.freq) == 1 || other.freq == 1) && other.count == 1 {
            return "YES"
        }
    }

    return "NO"
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}