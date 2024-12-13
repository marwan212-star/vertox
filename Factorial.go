package piscine

func Factorial(x int) int {
	if x > 20 || x < 0 {
		return 0
	}
	var res int = 1
	for i := 2; i <= x; i++ {
		res *= i
	}
	return res
}
