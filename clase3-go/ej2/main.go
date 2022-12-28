package main

import (
	"fmt"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Person   Person
	ID       int
	Position string
}

func (e Employee) PrintEmployee() {
	fmt.Printf("%+v \n", e)
}

func main() {
	persona := Person{
		ID:          1234,
		Name:        "Sebastian",
		DateOfBirth: "28/06/1999",
	}

	empleado := Employee{
		Person:   persona,
		ID:       4321,
		Position: "Software Developer",
	}

	empleado.PrintEmployee()

}
