package main

import (
	"errors"
	"fmt"
)

func main() {
	//  EJ 1
	fmt.Println("El salario sacando impuestos es: ", calcularImpuesto(100000))

	// EJ 2
	result, err := calcularPromedio(-1, 2, 3, 3, 3, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio de notas es: ", result)
	}

	//  EJ 3
	salario, err1 := calcularSalario(2000, "C")

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("El salario es: ", salario)
	}

	// EJ 4
	min := operation(minimum)
	max := operation(maximum)
	avg := operation(average)

	resultMin, errMin := min(1, 3, 4, 10, 12, 11, 9, 9, 5, 1)
	resultMax, errMax := max(1, 3, 4, 10, 12, 11, 9, 9, 5, 1)
	resultAvg, errAvg := avg(1, 3, 4, 10, 12, 11, 9, 9, 5, 1)

	if errAvg != nil {
		fmt.Println(errAvg)
	} else {
		fmt.Println("El promedio es : ", resultAvg)
	}

	if errMax != nil {
		fmt.Println(errMax)
	} else {
		fmt.Println("La nota maxima es : ", resultMax)
	}

	if errMin != nil {
		fmt.Println(errMin)
	} else {
		fmt.Println("La nota minima es : ", resultMin)
	}

	// EJ 5
	gato := animal(cat)
	perro := animal(dog)
	tigre := animal("tigre")

	amountGato, err1 := gato(15)
	amountPerro, err2 := perro(10)
	amountTigre, err3 := tigre(2)

	if err1 == nil {
		fmt.Println("La cantidad de alimento de gato a comprar es: ", amountGato)
	} else {
		fmt.Println(err1)
	}

	if err2 == nil {
		fmt.Println("La cantidad de alimento de perro a comprar es: ", amountPerro)
	} else {
		fmt.Println(err2)
	}

	if err3 == nil {
		fmt.Println("La cantidad de alimento de tigre a comprar es: ", amountTigre)
	} else {
		fmt.Println(err3)
	}

}

func calcularImpuesto(salario float64) float64 {
	if salario > 150000 {
		return salario - salario*0.27
	} else if salario > 50000 {
		return salario - salario*0.17
	} else {
		return salario
	}
}

func calcularPromedio(valores ...float64) (float64, error) {
	var sumaNotas float64
	valoresLength := float64(len(valores))
	for _, nota := range valores {
		if nota < 0 {
			return 0, errors.New("No se pueden introducir notas negativas.")
		}
		sumaNotas += nota
	}

	return sumaNotas / valoresLength, nil
}

func calcularSalario(minutosTrabajados float64, categoria string) (float64, error) {
	horasTrabajadas := minutosTrabajados / 60
	switch categoria {
	case "A":
		return (3000 * horasTrabajadas) * 1.5, nil
	case "B":
		return (1500 * horasTrabajadas) * 1.2, nil
	case "C":
		return 1000 * horasTrabajadas, nil
	default:
		return 0, errors.New("La categoria ingresada no existe.")
	}
}

func calcularMinimo(valores ...float64) (float64, error) {
	minimo := valores[0]
	for i := range valores {
		if valores[i] < 0 {
			return 0, errors.New("No se permiten calificaciones negativas.")
		} else if minimo > valores[i] {
			minimo = valores[i]
		}
	}
	return minimo, nil
}

func calcularMaximo(valores ...float64) (float64, error) {
	maximo := valores[0]
	for i := range valores {
		if valores[i] < 0 {
			return 0, errors.New("No se permiten calificaciones negativas.")
		} else if maximo < valores[i] {
			maximo = valores[i]
		}
	}
	return maximo, nil
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(estadistica string) func(valores ...float64) (float64, error) {
	switch estadistica {
	case minimum:
		return calcularMinimo
	case average:
		return calcularPromedio
	case maximum:
		return calcularMaximo
	}

	return nil
}

const (
	dog = "dog"
	cat = "cat"
)

func animalDog(cantidad float64) (float64, error) {
	if cantidad <= 0 {
		return 0, errors.New("La cantidad ingresada no puede ser menor a 0.")
	} else {
		return 10 * cantidad, nil
	}
}

func animalCat(cantidad float64) (float64, error) {
	if cantidad <= 0 {
		return 0, errors.New("La cantidad ingresada no puede ser menor a 0.")
	} else {
		return 5 * cantidad, nil
	}
}

func animalNotFound(cantidad float64) (float64, error) {
	return 0, errors.New("El animal ingresado no existe en el sistema.")
}

func animal(tipo string) func(cantidad float64) (float64, error) {
	switch tipo {
	case dog:
		return animalDog
	case cat:
		return animalCat
	default:
		return animalNotFound
	}
}
