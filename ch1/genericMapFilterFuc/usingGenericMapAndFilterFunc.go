package main

import "fmt"

func genericMap[T any, E any](input []T, f func(T) E) []E {
	result := make([]E, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

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

	input := []float64{-5.0, -2.0, 4.0, 8.0}
	result1 := genericMap(input, func(n float64) float64 {
		return n * n
	})
	fmt.Println(result1)

	greaterThanFive := MyFilter([]int{4, 6, 5, 2, 20, 1, 7}, func(i int) bool {
		return i > 5
	})
	fmt.Println(greaterThanFive)

	oddNumbers := MyFilter([]int{4, 6, 5, 2, 20, 1, 7}, func(i int) bool {
		return i%2 == 1
	})
	fmt.Println(oddNumbers)

	lengthGreaterThan3 := MyFilter([]string{"hello", "or", "the", "maybe"}, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println(lengthGreaterThan3)

}
