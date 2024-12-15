package piscine

func HoopCount(n int) string {
	var text = "Great, now move on to tricks"
	for i := 0; i <= n; i++ {
		if i >= 10 {
			return text
		}
	}
	return "Keep at it until you get it"
}
