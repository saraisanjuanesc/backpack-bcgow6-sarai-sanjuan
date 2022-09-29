package main

import "fmt"

func main() {
	palabra := "claves"
	fmt.Println("Longitud", len(palabra), "La palabra es: ")
	for i := 0; i < len(palabra); i++ {
		fmt.Println(palabra[i : i+1])
	}
}
