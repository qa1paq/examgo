package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error lenght argument")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("no such file or directory")
		return
	}
	var arr []float64

	ss := strings.Split(string(data), "\n")
	for _, w := range ss {
		a, err := strconv.Atoi(w)
		if err == nil {
			arr = append(arr, float64(a))
		}
	}

	arr = Bubblesort(arr)
	fmt.Printf("%s: %.0f\n", "Average", math.Round(Average(arr)))
	fmt.Printf("%s: %.0f\n", "Median", math.Round(Median(arr)))
	fmt.Printf("%s: %.0f\n", "Variance", math.Round(Variance(arr)))
	fmt.Printf("%s: %.0f\n", "Standard Deviation", math.Round(standardDeviation(arr)))
}

func Average(arr1 []float64) float64 {
	var sum float64

	for i := 0; i < len(arr1); i++ {
		sum = sum + arr1[i]
	}
	avg := sum / float64(len(arr1))
	return avg
}

func Bubblesort(arr1 []float64) []float64 {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1)-i-1; j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
	}
	return arr1
}

func Median(arr1 []float64) float64 {
	if len(arr1)%2 == 0 {
		a := len(arr1) / 2
		b := len(arr1)/2 - 1
		middle := (arr1[a] + arr1[b]) / 2
		return middle

	} else {
		middle := (len(arr1) + 1) / 2
		return float64(middle)
	}
}

func Variance(arr []float64) float64 {
	size := len(arr)
	sum := 0.0
	for _, key := range arr {
		sum = sum + key
	}
	res := float64(sum) / float64(size)

	sumRes := 0.0

	for _, key := range arr {
		sumRes = sumRes + (float64(key)-res)*(float64(key)-res)
	}
	result := float64(sumRes) / float64(size)
	return result
}

func standardDeviation(arr []float64) float64 {
	return (math.Sqrt(float64(Variance(arr))))
}
