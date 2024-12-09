package piscine

func Str_repit(str string, count int) string {
	if count > 99 {
		return "To Match Args"
	}
	val := ""
	for i := 0; i <= count; i++ {
		val += str
	}
	return val
}
