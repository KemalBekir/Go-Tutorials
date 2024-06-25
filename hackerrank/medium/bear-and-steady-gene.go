package main

func steadyGene(gene string) int32 {
    n := len(gene)
    exRep := n / 4 // expected repetition
    rep := map[rune]int{'A': 0, 'T': 0, 'G': 0, 'C': 0} // repetition map
    overRep := map[rune]int{'A': 0, 'T': 0, 'G': 0, 'C': 0} // over-repeated
    var pos []struct {
        letter rune
        index  int
    } // over-repeated position

    // check repetition
    for _, ch := range gene {
        rep[ch]++
    }
    overRep['A'] = max(rep['A']-exRep, 0)
    overRep['T'] = max(rep['T']-exRep, 0)
    overRep['G'] = max(rep['G']-exRep, 0)
    overRep['C'] = max(rep['C']-exRep, 0)

    if overRep['A'] == 0 && overRep['T'] == 0 && overRep['G'] == 0 && overRep['C'] == 0 {
        return 0
    }

    // Filter over-repeated
    for j, ch := range gene {
        if overRep[ch] > 0 {
            pos = append(pos, struct {
                letter rune
                index  int
            }{ch, j})
        }
    }

    var minLength int
    tempStart := 0
    letterCount := map[rune]int{'A': 0, 'T': 0, 'G': 0, 'C': 0}

    hasAllLetters := func() bool {
        return letterCount['A'] >= overRep['A'] &&
            letterCount['T'] >= overRep['T'] &&
            letterCount['G'] >= overRep['G'] &&
            letterCount['C'] >= overRep['C']
    }

    for k := 0; k < len(pos); k++ {
        letter, index := pos[k].letter, pos[k].index

        letterCount[letter]++
        for hasAllLetters() {
            startLetter, startIndex := pos[tempStart].letter, pos[tempStart].index
            newLength := index - startIndex + 1
            if minLength == 0 || newLength < minLength {
                minLength = newLength
            }
            letterCount[startLetter]--
            tempStart++
        }
    }

    return int32(minLength)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}