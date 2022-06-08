package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		a := 0
		var arr string
		arr1 := args[0]
		arr2 := args[1]
		for i := 0; i < len(arr1); i++ {
			for j := a; j < len(arr2); j++ {
				if arr1[i] == arr2[j] {

					arr = arr + string(arr1[i])

					j = len(arr2) - 1
				}
				a++
			}
		}
		if len(arr1) == len(arr) {
			fmt.Println(arr1)
		} else {
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}
