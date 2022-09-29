package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogFood(value int) float64 {
	return float64(value) * 10.0
}
func catFood(value int) float64 {
	return float64(value) * 5.0
}
func hamsterFood(value int) float64 {
	return float64(value) * .250
}
func tarantulaFood(value int) float64 {
	return float64(value) * .150
}
func animal(option string) (func(value int) float64, error) {
	switch option {
	case dog:
		return dogFood, nil
	case cat:
		return catFood, nil
	case hamster:
		return hamsterFood, nil
	case tarantula:
		return tarantulaFood, nil
	default:
		return nil, errors.New("El animal proporcionado No existe")
	}

}
func main() {
	animals, err := animal(dog)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(animals(8))
	}
}
