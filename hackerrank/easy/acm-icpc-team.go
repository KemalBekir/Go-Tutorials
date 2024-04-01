package main

func acmTeam(topic []string) []int32 {
	var maxTopics int32 = 0
	var maxTeams int32 = 0
	n := len(topic)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var knownTopics int32 = 0
			for k := 0; k < len(topic[i]); k++ {
				if topic[i][k] == '1' || topic[j][k] == '1' {
					knownTopics++
				}
			}
			if knownTopics > maxTopics {
				maxTopics = knownTopics
				maxTeams = 1
			} else if knownTopics == maxTopics {
				maxTeams++
			}
		}
	}
	return []int32{maxTopics, maxTeams}
}
