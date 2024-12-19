package piscine

func Repliction(str, str1, str2 string) string {
	//res := ""
	first := []rune(str)
	second := []rune(str1)
	thrid := []rune(str2)

	if len(second) > 1 && len(thrid) > 1 {
		return " "
	}
	for i := 0; i <= len(first)-1; i++ {
		if first[i] == second[0] {
			first[i] = thrid[0]
		}
	}
	//res += string(first)
	return string(first)
}
