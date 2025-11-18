package main

import "fmt"

type Person struct {
	Name string
}

type MyTypeMap interface {
	~int | ~float64 | ~string | Person | *Person // ~ cualqoer sub tipo de int, float64 o string
}

func genericMap[T MyTypeMap, E MyTypeMap](input []T, f func(T) E) []E {
	result := make([]E, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func main() {
	slice := []string{"go", "is", "fun"}
	result := genericMap(slice, func(i string) string {
		return i + i
	})
	fmt.Println(result)

	result2 := genericMap(slice, func(s string) int {
		return len(s)
	})
	fmt.Println(result2)

	slicePerson := []*Person{{Name: "Victor"}, {Name: "Felipe"}}
	resutl3 := genericMap(slicePerson, func(p *Person) Person {
		return Person{Name: p.Name + "!"}
	})
	fmt.Println(resutl3)

}
