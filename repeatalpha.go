package piscine

func RepeatAlpha(s string) string {
	res := ""
	if s == "" {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			count := int(s[i]) - 97
			for j := 0; j <= count; j++ {
				res += string(s[i])
			}

		} else if s[i] >= 'A' && s[i] <= 'Z' {
			count := int(s[i]) - 65
			for j := 0; j <= count; j++ {
				res += string(s[i])
			}

		} else {
			res += string(s[i])
		}
	}
	
	return res
}

/*ism := ""
if s == "" {
	return ""
}
for i := 0; i < len(s); i++ {
	if s[i] >= 'a' && s[i] <= 'z' {
		count := int(s[i]) - 97
		for j := 0; j < count; j++ {
			ism += string(s[i])
		}
	} else if s[i] >= 'A' && s[i] <= 'Z' {
		count := int(s[i]) - 65
		for j := 0; j < count; j++ {
			ism += string(s[i])
		}

	} else {
		ism += string(s[i])
	}
}
return ism*/
    