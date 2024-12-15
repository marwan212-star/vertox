package piscine

func LowarCase_Chaker(str string)string{
	var martix = []rune(str)
	for i := 0 ;i <= len(martix)-1;i++{
		if martix[i] >= 'A' && martix[i] <= 'Z'{
			martix[i] += 32
		}
	}
	return string(martix)
}