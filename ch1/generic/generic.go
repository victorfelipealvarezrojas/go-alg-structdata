package generic

import (
	"fmt"
)

type IStringer = interface {
	fnString() string
}

type AgeType int

func (i AgeType) fnString() string {
	return fmt.Sprintf("%d", i)
}

type NameType string

func (s NameType) fnString() string {
	return string(s)
}

type StudentType struct {
	Name string
	ID   int
	Age  float64
}

func (s StudentType) fnString() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}
func AddStudent[T IStringer](students []T, student T) []T {
	return append(students, student)
}
func main() {
	students := []NameType{}
	result := AddStudent(students, "Michael")
	result = AddStudent(result, "Jennifer")
	result = AddStudent(result, "Elaine")
	fmt.Println(result)

}
