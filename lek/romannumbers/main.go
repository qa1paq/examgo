package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var num int

	for _, w := range os.Args[1] {
		if w < 48 || w > 57 {
			fmt.Printf("ERROR: can not convert to roman digit\n")
			return
		}
		num = num*10 + int(w) - 48
	}
	if num >= 4000 {
		fmt.Printf("ERROR: can not convert to roman digit\n")
		return
	}
	res := ToRoman(num)
	fmt.Printf("%s\n", res)
}

func ToRoman(num int) string {
	num_array := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym_array := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	symCalc_arrray := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	var roman string
	var calcRoman string

	var i int
	for num > 0 {
		k := num / num_array[i]
		for j := 0; j < k; j++ {
			roman += sym_array[i]
			calcRoman += symCalc_arrray[i] + "+"
			num -= num_array[i]
		}
		i++
	}
	return calcRoman[:len(calcRoman)-1] + "\n" + roman
}
