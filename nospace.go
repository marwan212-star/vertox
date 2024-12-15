package piscine

func NoSpace(word string) string {
	var str string
	for i := 0 ; i <= len(word)-1;i++{
		if word[i] != ' '{
			str += string(word[i])
		}
	}
	return str
}
/*
	var s string
	for _, i := range word{
		if i != ' '{
			s += string(i)
		}
	}
*/