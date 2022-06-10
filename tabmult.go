package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	a := 0
	for _, r := range args[0] {
		if r >= '0' && r <= '9' {
			a = a*10 + int(r-'0')
		} else {
			return
		}
	}
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d", i, a, i*a)
		fmt.Println()
	}
}
