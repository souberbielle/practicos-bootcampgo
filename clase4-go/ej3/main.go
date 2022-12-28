package main

import (
	"errors"
	"fmt"
)

var salary int

var (
	customError = errors.New("Error: el salario es menor a 10.000")
)

func mySalary(value int) (string, error) {
	if value <= 10000 {
		return "", customError
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	salary, err := mySalary(9000)
	if errors.Is(err, customError) {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
