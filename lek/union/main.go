package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("\n")
		return
	}
	arg := os.Args[1] + os.Args[2]
	// arg := "zpadintonpaqefwtdjetyiytjneytjoeyjnejeyj"
	var res string
	res += string(arg[0])

	for i := 1; i < len(arg); i++ {
		k := 0
		for j := 0; j < len(res); j++ {

			if arg[i] == res[j] {
				k++
			}
			if j == len(res)-1 && k == 0 {
				res += string(arg[i])
			}
		}
	}
	fmt.Printf("%s\n", res)
}
