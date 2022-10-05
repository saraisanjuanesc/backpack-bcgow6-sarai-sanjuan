package main

import (
	"fmt"
	"math"
)

type product struct {
	Name   string
	Price  float64
	Amount int
}
type service struct {
	Name    string
	Price   float64
	Minutes int
}
type maintenance struct {
	Name  string
	Price float64
}

func totalProducts(prod *[]product, ch chan float64) {
	var total float64 = 0
	for _, value := range *prod {
		total += value.Price * float64(value.Amount)
	}
	ch <- total
}

func totalServices(serv *[]service, ch chan float64) {
	var total float64 = 0
	for _, value := range *serv {
		total += value.Price * (math.Ceil(float64(value.Minutes) / 30))
	}
	ch <- total
}

func totalMaintenance(maint *[]maintenance, ch chan float64) {
	var total float64 = 0
	for _, value := range *maint {
		total += value.Price
	}
	ch <- total
}

func main() {
	listProducts := []product{{"Producto1", 432.76, 4}, {"Producto2", 5224.77, 3}}
	listServices := []service{{"Servicio1", 34534.4, 24}, {"Servicio2", 3243.55, 45}}
	listMaintenance := []maintenance{{"Mantenimiento1", 64346.7}, {"Mantenimiento2", 23562.5}}

	canal := make(chan float64)
	go totalProducts(&listProducts, canal)
	go totalServices(&listServices, canal)
	go totalMaintenance(&listMaintenance, canal)
	var total float64 = 0
	for i := 0; i < 3; i++ {
		total += <-canal
	}

	fmt.Println("El Precio total de todo es: ", total)
}
