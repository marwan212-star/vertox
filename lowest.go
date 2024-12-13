package piscine

import (
	"strconv"
)
func Lowest(first,second int)string{
	if first > 10000 || first < 0 || second > 10000 || second < 0 {
		return "this is so match"
	}
	if first < second{
		return "The lowest number is : " + strconv.Itoa(first)
	}
	return "The lowest number is : " + strconv.Itoa(second)
}