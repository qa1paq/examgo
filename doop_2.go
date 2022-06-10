package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		var result int
		firstValue := Atoi(args[0])
		secondValue := Atoi(args[2])

		if args[1] == "*" {
			result = firstValue * secondValue
		} else if args[1] == "/" {
			result = firstValue / secondValue
		} else if args[1] == "+" {
			result = firstValue + secondValue
		} else if args[1] == "-" {
			result = firstValue - secondValue
		}
		fmt.Println(result)
	} else {
		fmt.Println()
	}
}

func Atoi(s string) int {
	n := 0
	runeStr := []rune(s)
	isPositive := 1
	i := 0
	for _, j := range runeStr {
		if j >= 0 || j <= 9 {
			if runeStr[0] == '-' {
				isPositive = -1
				i = 1
			}
			if runeStr[0] == '+' {
				i = 1
			}

			for ; i < len(runeStr); i++ {
				if runeStr[i] > 47 && runeStr[i] < 59 {
					n = n * 10
					n = n + int(rune(runeStr[i]-48))
				} else {
					return 0
				}
			}
			// if n < 9199999999999999999 && n > -9199999999999999999 {
			// 	return 0
			// }
			return n * isPositive
		}
	}
	return 0
}
