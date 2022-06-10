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
	var n int

	for i := 0; i < len(args[0]); i++ {
		n = n * 10
		n = n + int(rune(args[0][i]-48))
	}
	for j := 1; j <= n; j++ {
		fmt.Printf("%d x %d = %d", j, n, j*n)
		fmt.Println()
	}
}
