package main

import "fmt"

type user struct {
	Name     string
	LastName string
	Email    string
	Products []product
}
type product struct {
	Name   string
	Price  float64
	Amount int
}

func newProduct(name string, price float64) product {
	newProd := product{
		Name:  name,
		Price: price,
	}
	return newProd
}
func addProduct(us *user, prod *product, amount int) {
	var produccopy = *prod
	produccopy.Amount = amount
	us.Products = append(us.Products, produccopy)
}
func deleteProduct(us *user) {
	us.Products = nil
}

func main() {
	product1 := newProduct("Producto1", 57812.98)
	product2 := newProduct("Producto2", 57812.98)
	fmt.Println(product1)
	fmt.Println(product2)
	user1 := user{
		Name:     "Samantha",
		LastName: "Utrera",
		Email:    "samantha.utrera@gmail.com",
	}
	addProduct(&user1, &product1, 2)
	addProduct(&user1, &product2, 2)
	fmt.Println(user1)

	deleteProduct(&user1)
	fmt.Println(user1)

}
