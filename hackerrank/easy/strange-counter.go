package main

func strangeCounter(t int64) int64 {
	if t < 1 {
		return 0
	}
	var count int64 = 0
	var power int64 = 0
	base := int64(3)
	for count < t {
		base = 3 * (1 << power)
		if (base + count) < t {
			count += base
			power++
		} else {
			break
		}
	}
	if count == t {
		return 1
	}
	return base - (t - count - 1)
}
