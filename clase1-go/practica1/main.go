package main


import (
	"fmt"
	"practica1/clima"
)

var (
	name  = "Sebastian Souberbielle"
	address  = "Volteadores 1678"
)


func main() {
	fmt.Println(name + " " + address)
	fmt.Printf("Temperatura: %v °C Humedad: %v %% Presion: %v hPa\n", clima.Temperatura, clima.Humedad, clima.Presion)
}