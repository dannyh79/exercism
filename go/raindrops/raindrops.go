package raindrops

import "strconv"

func Convert(number int) string {
	var r string
	if number%3 == 0 {
		r += "Pling"
	}
	if number%5 == 0 {
		r += "Plang"
	}
	if number%7 == 0 {
		r += "Plong"
	}
	if len(r) == 0 {
		r = strconv.Itoa(number)
	}

	return r
}
