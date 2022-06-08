package main

import "fmt"

func main() {
	num := 17
	fmt.Println(ReverseBits(num))
}

func ReverseBits(num int) int {
	result := 0
	for i := 0; i < 8; i++ {
		result = result*2 + num%2
		num = num / 2
	}
	return result
}
