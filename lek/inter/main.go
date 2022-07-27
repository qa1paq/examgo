package main

import (
	"fmt"
)

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Printf("\n")
	// 	return
	// }
	arg1 := "abbcd"
	arg2 := "ujbs"
	var result string

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(arg2); j++ {
			if arg1[i] == arg2[j] {

				if len(result) > 0 {
					result = repeatChecker(result, string(arg1[i]))
				} else {
					result += string(arg1[i])
				}

				break
			}
		}
	}

	fmt.Printf("%s\n", result)
}

func repeatChecker(result string, letter string) string {
	checker := 0
	for k := 0; k < len(result); k++ {
		if result[k] == letter[0] {
			checker++
		}
		if k == len(result)-1 && checker == 0 {
			result += letter
			break
		}
	}
	return result
}
