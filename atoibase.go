package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, t string) int {
	res := 0
	lens := 0
	arr := map[rune]bool{}
	for _, c := range t {
		if arr[c] || c == '-' || c == '+' {
			return res
		}
		arr[c] = true
		lens++
	}
	if lens > 1 {
		for _, v := range s {
			count := 0
			if arr[v] {
				for _, j := range t {
					if j == v {
						break
					}
					count++
				}
				res = res*lens + count
			}
		}
	}
	return res
}
