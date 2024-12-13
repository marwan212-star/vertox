package piscine

import "strconv"

func Biggest(first, second int) string {
	if first > 10000 || first < 0 || second > 10000 || second < 0 {
		return "this is so match"
	}
	if first == second {
		return strconv.Itoa(first) + "this is =="

	}
	if first > second {
		return "The biggest number is : " + string(strconv.Itoa(first))
	}
	return "The biggest number is : " + strconv.Itoa(second)

}
