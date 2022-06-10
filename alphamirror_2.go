package main

import (
	"github.com/01-edu/z01"
)

func alphamirror(s string) {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			z01.PrintRune(65 + 90 - char)
		} else if char >= 'a' && char <= 'z' {
			z01.PrintRune(97 + 122 - char)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	alphamirror("abc3 xyz")
	// for i := 0; i < 255; i++ {
	// 	fmt.Print(i, " ")
	// 	z01.PrintRune(rune(i))
	// 	fmt.Println()
	// }
}
