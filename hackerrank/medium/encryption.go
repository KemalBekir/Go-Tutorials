package main

import (
	"strings"
	"math"
)

func encryption(s string) string {
    str := strings.ReplaceAll(s, " ", "")
    n := len(s)
    
    column := int32(math.Ceil(math.Sqrt(float64(n))))
    resultStr := ""
    hasSpace := false
    
    for i := int32(0); i < column; i++ {
        if hasSpace {
            resultStr += " "
        } else {
            hasSpace = true
        }
        
        for j := i; j < int32(n); j += column {
            resultStr += string(str[j])
        }
    }
    
    return resultStr
}