package main

import "fmt"

func main() {
	edad := 23
	empleado := true
	experiencia := 2
	sueldo := 101.000
	if edad > 22 {
		if empleado == true {
			if experiencia > 1 {
				if sueldo > 100.000 {
					fmt.Println("Puedes acceder al prestamo, y No se te cobrará interes")
				} else {
					fmt.Println("Puedes acceder al prestamo, pero se te cobrará interes")
				}
			} else {
				fmt.Println("No puedes acceder al prestamo, porque no tienes la experiencia establecida")
			}
		} else {
			fmt.Println("No puedes acceder al prestamo, porque no eres empleado")
		}

	} else {
		fmt.Println("No puedes acceder al prestamo, porque no tienes la edad")
	}
}
