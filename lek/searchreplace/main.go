package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	word := os.Args[1]
	let1 := os.Args[2]
	let2 := os.Args[3]

	for _, v := range word {
		if v == rune(let1[0]) {
			v = rune(let2[0])
		}
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
