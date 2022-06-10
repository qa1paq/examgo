package main

import "fmt"

func main() {
	var a byte = 0b01010100
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ReverseBits(a))
}

func ReverseBits(oct byte) byte {
	var res byte = 0
	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		fmt.Printf("!%08b\n", res)
		fmt.Printf("''%08b\n", oct)
		fmt.Printf("?%08b\n", oct&1)
		oct >>= 1
	}
	return res
}
