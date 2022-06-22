package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// arg := os.Args[1:]
	// if len(arg) != 2 {
	// 	return
	// }

	runes1 := "faya"
	runes2 := "fgvvfdxcacpolhyghbred"
	counter := 0
	k := 0
	for i := 0; i < len(runes1); i++ {
		for j := k; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				k = j + 1
				counter++
				break
			}
		}
	}

	if counter != len(runes1) {
		z01.PrintRune('\n')
	} else {
		for _, let := range runes1 {
			z01.PrintRune(let)
		}
		z01.PrintRune('\n')
	}
}
