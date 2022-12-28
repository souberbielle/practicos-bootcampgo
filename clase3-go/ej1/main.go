package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{ID: 1, Name: "Monitor", Price: 209.99, Description: "24 pulgadas", Category: "Tecnologia"},
	{ID: 2, Name: "Teclado RGB", Price: 99.99, Description: "Mecanico RGB", Category: "Tecnologia"},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

func getById(id int) (Product, error) {
	for _, product := range Products {
		if id == product.ID {
			return product, nil
		}
	}
	return Product{}, errors.New("Product not found.")
}

func main() {
	prod := Product{ID: 3, Name: "PC", Price: 499.99, Description: "MacBook Pro 15'", Category: "Tecnologia"}
	prod.Save()
	prod.GetAll()
	fmt.Println(getById(3))

	author := Author{}
}
