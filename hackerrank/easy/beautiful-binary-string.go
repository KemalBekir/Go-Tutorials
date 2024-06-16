package main

func beautifulBinaryString(b string) int32 {
    count := 0
    for i := 0; i < len(b)-2; i++ {
        if b[i:i+3] == "010" {
            count++
            b = b[:i] + "011" + b[i+3:]
            i += 2
        }
    }
    return int32(count)
}