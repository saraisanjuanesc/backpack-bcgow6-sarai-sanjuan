package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"], "")

	for key, element := range employees {
		if element > 21 {
			fmt.Println(key, "tiene más de 21 años")
		}
	}
	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
