package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}

	ism := 0
	for n != 0 {
		ism++
		n /= base
	}

	return ism
}
