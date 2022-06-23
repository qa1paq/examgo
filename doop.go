package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args[1:]

	if len(argument) == 3 {

		if !IsNumeric1(argument[0]) || !IsNumeric1(argument[2]) {
			return
		}

		a := atoi(argument[0])
		b := atoi(argument[2])
		if len(argument[0]) > 18 || len(argument[2]) > 18 {
			return
		} else if argument[1] == "/" && argument[2] == "0" {
			fmt.Println("No division by 0")
		} else if argument[1] == "%" && argument[2] == "0" {
			fmt.Println("No modulo by 0")
		} else if argument[1] == "+" {
			fmt.Println(a + b)
		} else if argument[1] == "-" {
			fmt.Println(a - b)
		} else if argument[1] == "*" {
			fmt.Println(a * b)
		} else if argument[1] == "/" {
			fmt.Println(a / b)
		} else if argument[1] == "%" {
			fmt.Println(a % b)
		}
	}
}

func atoi(b string) int {
	a := 0
	c := 1
	for _, word := range b {
		if word == '-' {
			c = -1
		}
		if word >= '0' && word <= '9' {
			a = a * 10
			a += (int(word) - '0')
		}
	}
	return a * c
}

func IsNumeric1(s string) bool {
	ff := 0
	for _, b := range s {
		if b >= '0' && b <= '9' || b == '-' {
			ff++
		}
	}
	if ff == len(s) {
		return true
	} else {
		return false
	}
}
