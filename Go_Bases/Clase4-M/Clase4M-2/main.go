package main

import (
	"errors"
	"fmt"
	"os"
)

func calcula(salary int) (err error) {
	if salary < 150_000 {
		err = errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return
}
func main() {
	var salary int = 100_000
	err := calcula(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe Pagar impuesto")
	}
}
