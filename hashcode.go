package piscine

func HashCode(dec string) string {
	var vaotrix string = ""
	for i := 0; i <= len(dec); i++ {
		hash_code := int(i+len(dec)) % 127
		if hash_code < 32 {
			hash_code += 33
		}
		vaotrix = vaotrix + string(rune(hash_code))
	}
	return vaotrix
}
