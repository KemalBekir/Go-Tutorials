package main

import "strings"

func gridSearch(G []string, P []string) string {
	rowLength := len(P[0])

	for i := 0; i < len(G); i++ {
		start := 0
		for {
			index := strings.Index(G[i][start:], P[0])
			if index == -1 {
				break
			}
			index += start

			match := true
			for j := 1; j < len(P); j++ {
				if (i + j) < len(G) {
					if strings.HasPrefix(G[i+j][index:], P[j]) && len(G[i+j]) >= index+rowLength {
						continue
					}
				}
				match = false
				break
			}

			if match {
				return "YES"
			}
			start = index + 1
		}
	}
	return "NO"
}
