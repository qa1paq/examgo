package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := len(os.Args) - 1
	var numbers []int
	for n != 0 {
		d := n % 10
		n /= 10
		numbers = append(numbers, d)
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		z01.PrintRune(rune(numbers[i] + 48))
	}
	z01.PrintRune('\n')
}
