package main

import "github.com/01-edu/z01"

func main() {
	runes := []rune("Hello World!\n")
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
}
