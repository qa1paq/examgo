package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	for _, v := range str {
		if ('a' <= v && v <= 'm') || ('A' <= v && v <= 'M') {
			v += 13
		} else if ('n' <= v && v <= 'z') || ('N' <= v && v <= 'Z') {
			v -= 13
		}
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
