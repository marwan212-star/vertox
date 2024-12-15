package piscine

func UpperCase_Chekar(str string) string {
	var gx = []rune(str)
	for i := 0; i <= len(gx)-1; i++ {
		if gx[i] >= 'a' && gx[i] <= 'z' {
			gx[i] -= 32
		}
	}
	return string(gx)
}
