package piscine

func CountAlpha(s string) int {
	var ism int
	for _, m := range s {
		if m >= 'A' && m <= 'z'{//|| m >= 'A' && m <= 'z'{
			ism++
		}
	}

	return ism
}

//	ism := 0
//	if ism == ' ' && s[i] > '0' || s[i] < '9' {
//			return 0
//		}
///	for i := 0 ; i < len(s) ; i++ {
//
//		ism++
//	}
//	return ism
//}
