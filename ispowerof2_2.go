package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		num := Atoi2(args[0])
		if num == 2 {
			fmt.Println("true")
			return
		}
		i := 0
		for i*i != num {
			i++
			if i*i == num {
				fmt.Println("true")
			} else if i*i > num {
				fmt.Println("false")
				return
			}
		}
	} else {
		fmt.Println()
	}
}

func Atoi2(s string) int {
	var n int
	runeStr := []rune(s)

	for i := 0; i < len(runeStr); i++ {
		n = n * 10
		n = n + int(rune(runeStr[i]-48))
	}
	return n
}
