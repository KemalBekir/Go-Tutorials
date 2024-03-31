package main

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
    black := int64(b)
    white := int64(w)
    blackCost := int64(bc)
    whiteCost := int64(wc)
    exchangeRage := int64(z)
    
    if blackCost + exchangeRage < whiteCost{
        return (blackCost + exchangeRage) * white + blackCost * black
    } else if whiteCost + exchangeRage < blackCost {
        return ( whiteCost + exchangeRage) * blackCost + whiteCost * white
    }else {
        return whiteCost* white + blackCost*black
    }
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
    // Convert to int64 immediately to avoid overflow in calculations
    b64 := int64(b)
    w64 := int64(w)
    bc64 := int64(bc)
    wc64 := int64(wc)
    z64 := int64(z)

    // Check if converting black to white or white to black is cheaper
    if bc64+z64 < wc64 {
        // Converting black to white is cheaper
        return (bc64+z64)*w64 + bc64*b64
    } else if wc64+z64 < bc64 {
        // Converting white to black is cheaper
        return (wc64+z64)*b64 + wc64*w64
    } else {
        // No conversion needed
        return wc64*w64 + bc64*b64
    }
}