package main

import (
	"fmt"
	"os"
)

type Product struct {
	ID     string
	Price  float64
	Number int
}

func main() {
	listProducts := []Product{{"P01", 86328.098, 8}, {"P02", 7633.98, 5}, {"P03", 82334.68, 3}, {"P04", 23456.99, 8}}
	listProducts = append(listProducts, Product{"P05", 3289047.23, 10})
	text := "ID;Precio;Cantidad\n"
	for _, value := range listProducts {
		text += fmt.Sprintf("%s;%.2f;%d\n", value.ID, value.Price, value.Number)
	}
	t := []byte(text)
	err := os.WriteFile("../inventorylist.csv", t, 0644)
	if err != nil {
		fmt.Println("ERROR \nCould not write to the file! ")
	} else {
		fmt.Println("File created")
	}
}
