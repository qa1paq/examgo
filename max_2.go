package main

func Max(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
