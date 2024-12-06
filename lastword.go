package piscine

func LastWord(s string) string {
	// ism := s
	// if len(ism) == 0 {
	// 	return "\n"
	// }

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' && s[i+1] != ' ' {
			return s[i+1:] + "\n"
		}
	}

	return  "\n"
}

// return ism[len(ism)-1] + "\n"

//if len(ism) > 0 {
//return ism[len(ism) - 1] + "\n"
//}

//}
//return "\n"

//ism := string(s)
//if len(ism) > 0 {
//return
//}
//return "\n"
