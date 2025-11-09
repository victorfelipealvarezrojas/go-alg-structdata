package main

import (
	"fmt"
)

type Student struct {
	Name string
	ID   int
	Age  float64
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}
func main() {
	students := []string{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")

	fmt.Println(result)
	students1 := []int{}
	result1 := addStudent(students1, 45)
	result1 = addStudent(result1, 64)
	result1 = addStudent(result1, 78)

	fmt.Println(result1)
	students2 := []Student{}
	result2 := addStudent(students2, Student{"John", 213, 17.5})
	result2 = addStudent(result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	fmt.Println(result2)
}
