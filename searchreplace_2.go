package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 || len(args[1]) != 1 || len(args[2]) != 1 {
		z01.PrintRune('\n')
		return
	}
	for i := 0; i < len(args[0]); i++ {
		if args[0][i] == args[1][0] {
			z01.PrintRune(rune(args[2][0]))
		} else {
			z01.PrintRune(rune(args[0][i]))
		}
	}
	z01.PrintRune('\n')
}
