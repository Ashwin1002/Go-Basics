package main

import "fmt"

func main() {

	var intSlice = []int{1, 2, 3, 4}
	fmt.Println(sumSlice(intSlice))

	var floatSlice = []float32{1.1, 2.2, 3.3, 4.4}
	fmt.Println(sumSlice(floatSlice))
}

// we can define which generics can be used here or we can pass `any` type
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, val := range slice {
		sum += val
	}
	return sum
}
