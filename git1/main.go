package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var set []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		// str := "189"
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Enter a valid number")
			return
		}
		set = append(set, num)
		// fmt.Println(set)

		FindBounds(set)
	}
}

func FindCenter(arr []int) int {
	le := len(arr)
	sort.Ints(arr)

	if le%2 == 1 {
		return (arr[(le-1)/2])
	} else {
		return (arr[le/2] + arr[le/2-1]) / 2
	}
}

func FindBounds(arr []int) {
	center := FindCenter(arr)
	lower := float64(center) * 0.72
	upper := float64(center) * 1.28

	fmt.Println(int(lower), int(upper))
}
