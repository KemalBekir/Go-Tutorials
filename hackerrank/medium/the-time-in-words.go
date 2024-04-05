package main

import (
	"fmt"
	"strings"
)

func hourToString(h int32) string {
	switch h {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "quarter"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	case 20:
		return "twenty"
	case 30:
		return "half"
	case 40:
		return "fourty"
	case 50:
		return "fifty"
	}

	return ""
}

func displayTime(timeParts ...string) string {
	rtp := []string{}
	for _, p := range timeParts {
		if p != "" {
			rtp = append(rtp, p)
		}
	}

	return strings.Join(rtp, " ")
}

func timeInWords(h int32, m int32) string {
	hour := hourToString(h)
	minutes := ""
	position := "past"
	entity := "minutes"

	if m == 0 {
		return fmt.Sprintf("%s o' clock", hour)
	}

	if m == 1 {
		entity = "minute"
	} else if m > 30 {
		hour = hourToString(h + 1)
		m = 60 - m
		position = "to"
	}

	if m%10 == 0 || m < 20 {
		minutes = hourToString(m)
	} else {
		min0 := hourToString(m / 10 * 10)
		min1 := hourToString(m % 10)
		minutes = fmt.Sprintf("%s %s", min0, min1)
	}

	if minutes == "quarter" || minutes == "half" {
		entity = ""
	}

	return displayTime(minutes, entity, position, hour)
}
