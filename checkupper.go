package piscine

func Check_Upper(str string) bool{
	var gx = true
	for i := 0; i <= len(str)-1;i++{
		if str[i] >= 'a'&& str[i] <= 'z'{
			return false
		}
	}
/*
	for _, c := range s{
		if s[c] >= 'a' && s[c] >= 'z'{
			return false
		}
	}
*/
	return gx
}