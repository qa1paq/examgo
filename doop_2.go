package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	a := (Atoi3(args[0]))
	b := (Atoi3(args[2]))
	if !IsNumeric(args[0]) || !IsNumeric(args[2]) {
		return
	}
	if len(args[0]) > 18 || len(args[2]) > 18 {
		fmt.Println(0)
		return
	}

	if args[1] == "/" && args[2] == "0" {
		fmt.Println("No division by 0")
	} else if args[1] == "%" && args[2] == "0" {
		fmt.Println("No modulo by 0")
	} else if args[1] == "+" {
		fmt.Println(a + b)
		return
	} else if args[1] == "*" {
		fmt.Println(a * b)
	} else if args[1] == "/" {
		fmt.Println(a / b)
	} else {
		return
	}
}

func IsNumeric(s string) bool {
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '-' {
			str = str + string(rune(s[i]))
		}
	}
	if str == s {
		return true
	}
	return false
}

func Atoi3(s string) int {
	sign := 1
	result := 0
	for _, r := range s {
		if r == '-' {
			sign = -1
		}
		if r >= '0' && r <= '9' {
			result = result * 10
			result = result + int(rune(r)-48)
		}
	}
	return result * sign
}
