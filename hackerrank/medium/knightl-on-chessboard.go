package main

import "container/list"

type Point struct {
    x, y, dist int32
}

func bfs(n, a, b int32) int32 {
    directions := [][2]int32{{a, b}, {a, -b}, {-a, b}, {-a, -b},
                             {b, a}, {b, -a}, {-b, a}, {-b, -a}}
    queue := list.New()
    queue.PushBack(Point{0, 0, 0})
    visited := make(map[[2]int32]bool)
    visited[[2]int32{0, 0}] = true

    for queue.Len() > 0 {
        current := queue.Remove(queue.Front()).(Point)
        if current.x == n-1 && current.y == n-1 {
            return current.dist
        }

        for _, dir := range directions {
            nx, ny := current.x + dir[0], current.y + dir[1]
            if nx >= 0 && nx < n && ny >= 0 && ny < n {
                if !visited[[2]int32{nx, ny}] {
                    visited[[2]int32{nx, ny}] = true
                    queue.PushBack(Point{nx, ny, current.dist + 1})
                }
            }
        }
    }

    return -1
}

func knightlOnAChessboard(n int32) [][]int32 {
    result := make([][]int32, n-1)
    for i := int32(0); i < n-1; i++ {
        result[i] = make([]int32, n-1)
        for j := int32(0); j < n-1; j++ {
            result[i][j] = bfs(n, i+1, j+1)
        }
    }
    return result
}