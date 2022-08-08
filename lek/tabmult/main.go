package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var num int
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	for i := 1; i <= 9; i++ {
		num = i * arg
		Println(Itoa(i) + " x " + Itoa(arg) + " = " + Itoa(num))
	}
}

func Itoa(n int) string {
	var ch string

	if n < 0 {
		n = -n
		ch = "-"
	}
 
	digits := "0123456789"

	if n < 10 {
		return ch + digits[n:n+1]
	}

	return ch + Itoa(n/10) + digits[n%10:n%10+1]
}

func Println(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
