package main

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	black := int64(b)
	white := int64(w)
	blackCost := int64(bc)
	whiteCost := int64(wc)
	exchangeRage := int64(z)

	if blackCost+exchangeRage < whiteCost {
		return (blackCost+exchangeRage)*white + blackCost*black
	} else if whiteCost+exchangeRage < blackCost {
		return (whiteCost+exchangeRage)*blackCost + whiteCost*white
	} else {
		return whiteCost*white + blackCost*black
	}
}
