package piscine

func CountChar(str string, c rune) int {
	var ism int
	for _, k := range str {
		if k == c {
			ism++
		}
	}
	return ism
}
