package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var a, b, c string
	error := "No vowels"
	piglatin := "ay"
	if len(args) == 1 {
		for i := 0; i < len(args[0]); i++ {
			if args[0][i] == 'a' || args[0][i] == 'e' || args[0][i] == 'y' || args[0][i] == 'u' || args[0][i] == 'i' || args[0][i] == 'o' || args[0][i] == 'A' || args[0][i] == 'E' || args[0][i] == 'Y' || args[0][i] == 'U' || args[0][i] == 'I' || args[0][i] == 'O' {
				a = a + string(args[0][i])
				for l := 0; l < i; l++ {
					c = c + string(args[0][l])
				}
				i++
				for j := i; j < len(args[0]); j++ {
					b = b + string(args[0][j])
				}

			}
		}
		if a == "" {
			fmt.Println(error)
			return
		}
		result := a + b + c + piglatin
		fmt.Println(result)
	}
	if len(args) > 1 {
		z01.PrintRune('\n')
	}
}
