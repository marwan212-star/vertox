package piscine

import "strconv"

func CountSheep(num int) string {
	if num == 0 {
		return ""
	}
	if num < 0 {
		return ""
	}
	var count string
	for i := 1; i <= num; i++ {
		count += strconv.Itoa(i) + " Sheep..."
	}
  return count
}
