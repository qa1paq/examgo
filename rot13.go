package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		for i := 0; i < len(args[0]); i++ {
			if (args[0][i] >= 'A' && args[0][i] < 'N') || (args[0][i] >= 'a' && args[0][i] < 'n') {
				z01.PrintRune(rune(args[0][i]) + 13)
			} else if args[0][i] >= 'N' && args[0][i] <= 'Z' || args[0][i] >= 'n' && args[0][i] <= 'z' {
				z01.PrintRune(rune(args[0][i]) - 13)
			} else {
				z01.PrintRune(rune(args[0][i]))
			}
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}
