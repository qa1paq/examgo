package main

import "fmt"

func main() {
	fmt.Println(Swapbits(2))
}

func Swapbits(a byte) byte {
	return a<<4 | a>>4
}
