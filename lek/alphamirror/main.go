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
	for _, w := range arg {
		if w >= 'a' && w <= 'z' {
			w = 'z' - w + 'a'
		} else if w >= 'A' && w <= 'Z' {
			w = 'Z' - w + 'A'
		}
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
