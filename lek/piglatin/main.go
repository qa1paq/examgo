package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	vowels := "ioueaIOUEA"
	isVowel := false
	arg := os.Args[1]

	for _, w := range vowels {
		if rune(arg[0]) == w {
			arg += "ay"
			isVowel = true
		}
	}

	if isVowel {
		Println(arg)
	} else {
		Change(arg, vowels)
	}
}

func Change(arg, vowels string) {
	temp := ""
	for _, w := range arg {
		temp += string(w)
		for _, e := range vowels {
			if w == e {
				arg = arg[len(temp)-1:] + arg[:len(temp)-1] + "ay"
				Println(arg)
				return
			}
		}
	}
	Println("No vowels")
}

func Println(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
