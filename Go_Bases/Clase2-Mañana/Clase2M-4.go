package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minimumOperation(values ...float64) float64 {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}
func averageOperation(values ...float64) float64 {
	var sum float64
	for _, value := range values {
		sum = sum + value
	}
	return sum / float64(len(values))
}

func maximumOperation(values ...float64) float64 {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func operation(option string) (func(values ...float64) float64, error) {
	switch option {
	case minimum:
		return minimumOperation, nil
	case average:
		return averageOperation, nil
	case maximum:
		return maximumOperation, nil
	default:
		return nil, errors.New("La opción proporcionada No existe")
	}

}
func main() {
	minFunc, err1 := operation(minimum)
	averageFunc, err2 := operation(average)
	maxFunc, err3 := operation(maximum)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("La calificación mínima es: ", minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("El promedio de las calificaciones es: ", averageFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("La calificación máxima es: ", maxFunc(2, 3, 3, 4, 10, 2, 4, 5))
	}

}
