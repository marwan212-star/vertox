package piscine

func FishAndChips(n int) string {
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}else if n%2 == 0 {
		return "fich"
	} else if n%3 == 0 {
		return "chips"
	} else if n < 0 {
		return "error: number is negative"
	} else {
		return "error: non divisible"
	}
}
