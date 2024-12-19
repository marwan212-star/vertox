package piscine

func Atoi(str string) int {
	i := 0
	r := 0
	min := 1
	if str[0] == '-' {
		min = -1// handel nigtev number
		i++
	}
	for ; i < len(str); i++ {
		r *= 10
		r += int(str[i] - 48)
	}
	return r * min
}
