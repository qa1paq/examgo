package main

import "github.com/01-edu/z01"

func main() {
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		count++
		if count%2 == 0 {
			z01.PrintRune(i - ' ')
		} else {
			z01.PrintRune(i)
		}

	}
	z01.PrintRune('\n')
}
