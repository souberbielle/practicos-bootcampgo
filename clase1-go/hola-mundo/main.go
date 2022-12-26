package main


import (
	"fmt"
	"hola-mundo/numbers"
)

func main() {
	fmt.Println("Hola Mundo!")
	myPi  := numbers.Pi
	fmt.Println(myPi)
	fmt.Printf("tipo: %T\n", myPi)
}