package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	res := Atoi(os.Args[1])
	if IsPowerOf2(res) == 0 {
		PrintRune("true")
	} else {
		PrintRune("false")
	}
}

func Atoi(str string) int {
	var n int
	for _, w := range str {
		n = n*10 + int(w) - 48
	}
	return n
}

func IsPowerOf2(n int) int {
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

func PrintRune(str string) {
	for _, w := range str {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
