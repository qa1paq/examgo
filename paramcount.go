package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	count := 0
	for i := 0; i < len(args); i++ {
		count++
	}
	fmt.Println(count)
}
