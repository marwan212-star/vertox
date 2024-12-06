package piscine


func CountIf(f func(string) bool, tab []string) int {
	a := 0
	for _, iss := range tab {
		if f(iss) {
			a++
		}
	}
	return a
}