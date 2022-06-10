package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = f(res, a[i])
	}
	num := strconv.Itoa(res)
	PrintRune(num)
	z01.PrintRune('\n')
}

func PrintRune(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
}
