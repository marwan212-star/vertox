package piscine

// func CamelToSnakeCase(s string) string {
// ism := string(s[0])
// if ism[0] > 'A' || ism[0] < 'Z' {
// 	if ism[0] > 'a' || ism[0] < 'z' {
// 		return s
// 	}
// }
// m := s[len(s)-1]
// if m >= 'A' && m <= 'Z' {
// 	return s
// }
// for i := 0; i < len(s); i++ {
// 	if s[i] >= 'A' && s[i] <= 'Z' {
// 		if s[i-1] >= 'A' && s[i-1] <= 'Z' {
// 			return s
// 		}
// 		ism += "_"
// 	}
// 	ism += string(s[i])
// }

// return ism

/*k := string(s[0])
if k[0] < 'A' || k[0] > 'Z' {
	if k[0] < 'a' || k[0] > 'z' {
		return s
	}
}
a := s[len(s)-1]
if a >= 'A' && a <= 'Z' {
	return s
}

for i := 1; i < len(s); i++ {

	if s[i] >= 'A' && s[i] <= 'Z' {
		if s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}
		k += "_"
	}
	k += string(s[i])

}
return k*/
// }

func check(s string) bool {
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return false
	}
	
	for i := 0 ; i < len(s) -1  ; i++{
		if s[i] >= '0' && s[i] <= '9' {
			return false
		}
		if s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' &&   s[i+1] <= 'Z' {
			return false
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if !check(s) {
		return s
	}
	
	ism := string(s[0])
	for i := 1; i < len(s); i++ {
		if s[i-1] >= 'a' && s[i-1] <= 'z'&& s[i] >= 'A' && s[i] <= 'Z' {
			ism += "_"
			ism += string(s[i])
		} else {
			ism += string(s[i])
		}
	}

	return ism
}
