package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	arr1 := []rune{}
	count := 0
	to := 0
	if len(args) == 1 {
		for i := len(args[0]) - 1; i >= 0; i-- {
			arr1 = append(arr1, rune(args[0][i]))
		}
		for _, r := range arr1 {
			if r != ' ' {
				break
			}
			count++
		}
		for i := count; i < len(arr1); i++ {
			if arr1[i] == ' ' {
				to = i
				break
			}
		}

		for i := to - 1; i >= count; i-- {
			z01.PrintRune(arr1[i])
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
