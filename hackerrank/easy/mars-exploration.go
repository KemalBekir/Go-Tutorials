package main

func marsExploration(s string) int32 {
    var cs int32 = 0
    
    for i := 0; i < len(s); i += 3 {
        if s[i] != 'S' {
            cs++
        }
    }
    
    for i := 1; i < len(s); i += 3 {
        if s[i] != 'O' {
            cs++
        }
    }
    
    for i := 2; i < len(s); i += 3 {
        if s[i] != 'S' {
            cs++
        }
    }
    
    return cs
}