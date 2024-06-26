package main

func morganAndString(a string, b string) string {
    a += "z"
    b += "z"
    i, j := 0, 0
    var res []byte

    for i < len(a) && j < len(b) {
        if a[i:] < b[j:] {
            res = append(res, a[i])
            i++
        } else {
            res = append(res, b[j])
            j++
        }
    }