package main

import (
	"fmt"
	"os"
)

func calcula(salary int) (err error) {
	if salary < 150_000 {
		err = fmt.Errorf("error: el mínimo imponible s de: $15000 y el salio ingresado es de $%d", salary)
	}
	return
}

func main() {
	var salary int = 150_000
	err := calcula(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe Pagar impuesto")
	}
}
