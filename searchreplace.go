package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 || len(args[1]) != 1 || len(args[2]) != 1 {
		fmt.Println()
		return
	}

	var arr []rune
	for _, r := range args[0] {
		if rune(args[1][0]) == r {
			arr = append(arr, rune(args[2][0]))
		} else {
			arr = append(arr, r)
		}
	}
	fmt.Println(string(arr))

}
