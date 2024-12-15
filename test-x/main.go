package main

import "fmt"

//import "x_x"

func main() {

	fmt.Println(Say_Hello("Jack Pell", 1))

}

func Say_Hello(name string, number int) string {

	if number < 0 || number > 1 {
		return " "
	}

	if number == 1 {
		return "Hello " + Cap(name)
	} else if number == 0 {
		return "Bye " + Cap(name)
	}
	return "Bye " + Cap(name)
}

func Cap(names string) string {
	runes := []rune(names)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			 runes[i]-= 32
		} else if runes[i] >= 'A' && runes[i] >= 'Z' {
			continue
		}
	}
	return string(runes)
}
