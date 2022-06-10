package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		return
	}
	for _, let := range arg[0] {
		z01.PrintRune(let)
	}
	z01.PrintRune('\n')
}
