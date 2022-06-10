package main

import (
	"fmt"
)

func main() {
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := Max(arrInt)
	fmt.Println(max)
}

func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
