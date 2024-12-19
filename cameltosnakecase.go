package piscine

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
