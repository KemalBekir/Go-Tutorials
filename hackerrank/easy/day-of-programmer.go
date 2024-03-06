package main

import "strconv"

func dayOfProgrammer(year int32) string {
	if ((year%400 == 0) || (year%4 == 0 && year%100 != 0)) && year > 1918 || (year%4 == 0 && year < 1918) {
		return "12.09." + strconv.Itoa(int(year))
	} else if year == 1918 {
		return "26.09.1918"
	} else {
		return "13.09." + strconv.Itoa((int(year)))
	}
}
