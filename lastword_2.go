package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1]
	// arg := "FOR PONY"
	runes := []rune(arg)
	counter := 0
	if len(arg) > 1 {
		for i := len(runes) - 1; i >= 0; i-- {
			if runes[i] == ' ' {
				counter++
			} else {
				break
			}
		}
		var str string
		for i := len(runes) - counter - 1; i >= 0; i-- {
			if runes[i] != ' ' {
				str += string(runes[i])
			} else {
				break
			}
		}
		for i := len(str) - 1; i >= 0; i-- {
			z01.PrintRune(rune(str[i]))
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}
