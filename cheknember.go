package piscine

func CheckNumber(arg string) bool {
	for _, iss := range arg {
		if iss >= '0' && iss <= '9' {
			return true
		}
	}
	return false
}
