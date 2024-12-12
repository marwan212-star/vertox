package piscine

func Zero_len(str string) int {
	var count int
	for _, i := range str {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i >= '1' && i <= '9' || i == ' ' || i >= 32 && i <= 46 {
			count++
		}
	}
	return count
}
