package main

import "fmt"

func main() {
	// Chunk([]int{}, 10)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	var arr [][]int
	var buf []int
	k := 0
	for i := 0; i < len(slice); i++ {
		buf = append(buf, slice[i])
		k++
		if k == size {
			arr = append(arr, buf)
			buf = nil
			k = 0
		}
	}
	if k != 0 {
		arr = append(arr, buf)
	}
	fmt.Println(arr)
}
