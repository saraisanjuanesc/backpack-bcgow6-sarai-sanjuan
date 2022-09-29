package main

import (
	"errors"
	"fmt"
)

func monthSalary(hours int, salary float64) (result float64, err error) {
	if hours < 0 || hours < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	result = float64(hours) * salary
	if result >= 150_000 {
		result -= result * 0.1
	}
	return
}

func calculateBonus(salaries []float64) (result float64, err error) {
	var bestSalary float64
	numberMonths := len(salaries)
	for _, salary := range salaries {
		if salary < 0 {
			err = errors.New("error: No puede haber salarios negativos")
			return
		}
		if salary > bestSalary {
			bestSalary = salary
		}
	}
	result = bestSalary / 12 * float64(numberMonths)
	return
}

func main() {
	var hours int = 120
	var salary float64 = 150_000.00

	salaryM, err := monthSalary(hours, salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario mensual es: $%.2f\nHabiendo trabajado un total de: %d horas y ganando $%.2f por horan\n", salaryM, hours, salary)
	}

	salaries := []float64{150000, 300000, 400000, 230000, 130000}
	bonus, errbonus := calculateBonus(salaries)
	if errbonus != nil {
		fmt.Println(errbonus)
	} else {
		fmt.Printf("\nEl agunaldo del empleado es: $%.2f\n", bonus)
	}
}
