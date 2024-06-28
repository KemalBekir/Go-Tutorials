package main

func alternate(s string) int32 {
    strSet := make(map[rune]struct{})
    for _, ch := range s {
        strSet[ch] = struct{}{}
    }

    var maxRes int32 = 0
    for ch1 := range strSet {
        for ch2 := range strSet {
            if ch1 != ch2 {
                // Create the string with only characters ch1 and ch2
                t := make([]rune, 0)
                for _, ch := range s {
                    if ch == ch1 || ch == ch2 {
                        t = append(t, ch)
                    }
                }

                // Validate the created string
                if validate(string(t)) {
                    if int32(len(t)) > maxRes {
                        maxRes = int32(len(t))
                    }
                }
            }
        }
    }
    return maxRes
}
