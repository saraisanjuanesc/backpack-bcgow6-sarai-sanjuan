package main

import "fmt"

func impuesto(salario float64) (result float64) {

	if salario > 50000 {
		result = salario * 0.17
	}
	if salario > 150000 {
		result += salario * 0.10
	}
	return
}
func main() {
	var salario float64 = 60000
	fmt.Println(impuesto(salario))
}
