package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position int
	Person
}

func (s *Employee) PrintEmployee() {
	fmt.Println(fmt.Sprintf("%d-%d-%s", s.ID, s.Position, s.Person.Name))
}

func main() {
	person := Person{
		ID:          1,
		Name:        "Gabriel",
		DateOfBirth: "10/10/1992",
	}

	employee := Employee{
		ID:       1,
		Position: 1,
		Person:   person,
	}

	employee.PrintEmployee()
}
