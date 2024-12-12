package piscine

import "strconv"

func FromTo(from int, to int) string {
	if to > 99 || to < 0 && from < 0 || from > 99 {
		return "Invalid\n"
	}
	if from == to {
		return strconv.Itoa(from) + "\n"
	}
	var num string
	for x := from; x <= to; x++ {
		if x > 0 && x < 10 {
			num += "0" + strconv.Itoa(x) + ", "
			if x >= 10 {
				num += strconv.Itoa(x)
			}
		}
	}
	return num
}
