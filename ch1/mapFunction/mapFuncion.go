package main

import "fmt"

func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func main() {
	slice := []int{1, 5, 2, 7, 4}
	result := MyMap(slice, func(i int) int {
		return i * i
	})
	fmt.Println(result)
}
