package main

import (
	"errors"
	"fmt"
)

func calculate_salary(minutes int, category string) (salary float64, err error) {
	if minutes < 0 {
		err = errors.New("Los minutos no pueden ser Negativos")
		return
	}
	if category == "C" {
		salary = (float64(minutes) / 60) * 1000.00
	} else if category == "B" {
		salary = ((float64(minutes) / 60) * 1500.00)
		salary += salary * 0.2
	} else if category == "A" {
		salary = ((float64(minutes) / 60) * 3000.00)
		salary += salary * 0.5
	}
	return
}
func main() {
	result, err := calculate_salary(6000, "B")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es: ", result)
	}
}
