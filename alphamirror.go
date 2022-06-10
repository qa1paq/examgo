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
	var result []rune
	for _, r := range args[0] {
		if r >= 'a' && r <= 'z' {
			a := 'z' - r + 'a'
			result = append(result, a)
		} else if r >= 'A' && r <= 'Z' {
			a := 'Z' + 'A' - r
			result = append(result, a)
		} else {
			result = append(result, r)
		}
	}
	fmt.Println(string(result))
}
