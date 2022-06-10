package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(11))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	y := false
	for nb >= 2 {
		// fmt.Println(nb)
		for i := 2; i < nb; i++ {
			if nb%i != 0 {
				y = true
			} else {
				y = false
				break
			}
		}
		if y == false {
			nb--
		} else {
			return nb
		}
	}
	return 0
}
