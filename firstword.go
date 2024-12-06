package piscine

// import "strings"

// import "fmt"

func FirstWord(s string) string {
	n := 0
	for len(s) > n && s[n] == ' ' {
		n++
	}
	start := n
	for len(s) > n && s[n] != ' ' {
		n++
	}

	return s[start:n] + "\n"
}
