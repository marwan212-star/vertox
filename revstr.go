package piscine

func Rev_Str(str string) string {
	var xg = ""
	var gx = []rune(str)
	for i := len(gx) - 1; i >= 0; i-- {
		xg += string(gx[i])
	}
	return xg
}
