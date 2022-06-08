package main

import (
	"fmt"
	"os"
)

func main() {
	ss := os.Args[1:]
	// ss := "faya"
	// ss1 := "fgvvfdxcacpolhyghbreda"
	r1 := ss[0]
	r2 := ss[1]
	var result string
	counter := 0
	for i := 0; i < len(r1); i++ {
		for j := counter; j < len(r2); j++ {
			if r1[i] == r2[j] {
				result = result + string(r1[i])
				j = j + len(r2) - 1
			}
			counter++
		}
	}
	if result == r1 {
		fmt.Println(result)
	}
}
