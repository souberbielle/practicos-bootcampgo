package main

import "fmt"

var salary int

type customError struct{}

func (e *customError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func mySalary(value int) (string, *customError) {
	if value < 150000 {
		return "", &customError{}
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	salary, err := mySalary(180000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}

	salary1, err := mySalary(130000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary1)
	}
}
