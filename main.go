package main

import (
	"algstruc/ch1/generic"
	"fmt"
)

func main() {
	students := []generic.NameType{}
	result := generic.AddStudent(students, "Michael")
	result = generic.AddStudent(result, "Jennifer")
	result = generic.AddStudent(result, "Elaine")
	fmt.Println(result)

	students1 := []generic.AgeType{}
	result1 := generic.AddStudent(students1, 45)
	result1 = generic.AddStudent(result1, 64)
	result1 = generic.AddStudent(result1, 78)
	fmt.Println(result1)

	students2 := []generic.StudentType{}
	result2 := generic.AddStudent(students2, generic.StudentType{Name: "John", ID: 213, Age: 17.5})
	result2 = generic.AddStudent(result2, generic.StudentType{Name: "James", ID: 111, Age: 18.75})
	result2 = generic.AddStudent(result2, generic.StudentType{Name: "Marsha", ID: 110, Age: 16.25})

	fmt.Println(result2)
}
