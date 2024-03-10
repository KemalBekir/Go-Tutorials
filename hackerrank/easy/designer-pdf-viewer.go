package main

func designerPdfViewer(h []int32, word string) int32 {
	var maxHeight int32
	for _, v := range word {
		if h[v-97] > maxHeight {
			maxHeight = h[v-97]
		}
	}
	return maxHeight * int32(len(word))
}
