package main

import (
	"errors"
	"fmt"
)

var salary int

type customError struct{}

func (e *customError) Error() string {
	return "Error: el salario es menor a 10.000"
}

func mySalary(value int) (string, *customError) {
	if value <= 10000 {
		return "", &customError{}
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	salary, err := mySalary(9000)
	var errorLowSalary customError
	if errors.Is(err, &errorLowSalary) {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
