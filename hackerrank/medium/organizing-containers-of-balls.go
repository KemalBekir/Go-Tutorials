package main

func organizingContainers(container [][]int32) string {
    n := len(container)
    row := make([]int32, n)
    col := make([]int32, n)
    
    for i:= 0; i < n; i++{
        for j := 0; j < n; j++ {
            row[i] += container[i][j]
            col[i] += container[j][i]
        }
    }
    
    sortSlice := func(slice []int32) {
        sort.Slice(slice, func(i, j int) bool {
            return slice[i] < slice[j]
        })
    }
    
    sortSlice(row)
    sortSlice(col)
    
    for i := 0; i < n; i++ {
        if row[i] != col[i] {
            return "Impossible"
        }
    }
    
    return "Possible"
}