package piscine

func Itoa(n int) string {
	a := ""
	sign := ""
	if n < 0 {
		sign = "-" // Handle negative numbers
		n = -n
	}

	for n > 0 {
		digit := n % 10
		a = string(rune('0'+digit)) + a // Convert digit to character
		n = n / 10
	}

	return sign + a
}
