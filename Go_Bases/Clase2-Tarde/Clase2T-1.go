package main

import "fmt"

type alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a alumno) detalle() {
	fmt.Println("Nombre: ", a.Nombre, "\nApellido: ", a.Apellido, "\nDNI: ", a.DNI, "\nFecha: ", a.Fecha)
}
func main() {
	a1 := alumno{"Marcos", "Marquez", "BA000589", "15/09/2022"}
	a1.detalle()
}
