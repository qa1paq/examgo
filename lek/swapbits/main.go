package main

import "fmt"

func main() {
	var b byte = 4
	fmt.Println(SwapBits(b))
}

func SwapBits(octet byte) byte {
	return octet>>4 | octet<<4
}
