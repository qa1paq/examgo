package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(1))
	fmt.Println(FindPrevPrime(2))
}

func FindPrevPrime(nb int) int {
	if nb == 2 {
		return 2
	} else if nb <= 1 {
		return 0
	}

	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return FindPrevPrime(nb - 1)
		}
	}
	return nb
}
