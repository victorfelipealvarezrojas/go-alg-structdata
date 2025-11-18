package main

import "fmt"

func MyFilter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, value := range input {
		if f(value) == true {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	res := MyFilter(input, func(num float64) bool {
		if num <= 10.0 {
			return true
		}
		return false
	})

	fmt.Println(res)
}
