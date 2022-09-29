package main

import (
	"errors"
	"fmt"
)

func average(values ...float64) (result float64, err error) {
	var sum float64
	for _, value := range values {
		if value < 0 {
			err = errors.New("La calificación No puede ser Negativa")
			return
		}
		sum = sum + value
	}
	result = sum / float64(len(values))
	return

}
func main() {
	result, err := average(2, 4, 6, -9, 8, 9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
