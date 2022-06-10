package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	str := ""
	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] != ' ' {
			str = string(arg[i]) + str
		} else if len(str) != 0 {
			break
		}
	}
	PrintRune(str)
	z01.PrintRune('\n')
}

func PrintRune(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
}
