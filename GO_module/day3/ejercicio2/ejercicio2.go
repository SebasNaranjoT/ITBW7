package main

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Person
	WorkId   int
	Position string
}

func (e Employee) PrintEmployee() {
	fmt.Printf("%+v", e)
}

func main() {
/* 	p := Person{
		Id:          1234,
		Name:        "sebastian",
		DateOfBirth: "06/03/1998",
	}
 */
	e := Employee{
		Person: Person{
			Id:          1234,
			Name:        "sebastian",
			DateOfBirth: "06/03/1998",
		},
		WorkId:   82345,
		Position: "Senior engineer",
	}

	e.PrintEmployee()
}
