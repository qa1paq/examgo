package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println()
		return
	}
	d := 1
	c := 1
	a := 0
	b := 0
	for i, r := range args[0] {
		if r >= '0' && r <= '9' {
			a = a*10 + int(r-'0')
		} else if i == 0 && r == '-' {
			c = -1
		} else {
			fmt.Println("0")
			return
		}
	}
	for i, r := range args[2] {
		if r >= '0' && r <= '9' {
			b = b*10 + int(r-'0')
		} else if i == 0 && r == '-' {
			d = -1
		} else {
			fmt.Println("0")
			return
		}
	}
	b = b * d
	a = a * c
	if args[1][0] == '+' {
		fmt.Println(a + b)
	} else if args[1][0] == '-' {
		fmt.Println(a - b)
	} else if args[1][0] == '*' {
		fmt.Println(a * b)
	} else if args[1][0] == '/' {
		if b == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(a / b)
		}
	} else if args[1][0] == '%' {
		if b == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(a % b)
		}
	} else {
		fmt.Println("0")
		return
	}
}
