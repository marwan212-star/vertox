package piscine

func RetainFirstHalf(str string) string {
	if len(str) == 1 {
		return str
	}
	var devsion int = len(str) / 2
	return str[:devsion]
}
