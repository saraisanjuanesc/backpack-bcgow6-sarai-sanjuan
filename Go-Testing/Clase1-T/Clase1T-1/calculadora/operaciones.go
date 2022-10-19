package calculadora

import (
	"fmt"
	"sort"
)

func Sumar(num1, num2 int) int {
	return num1 + num2
}

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Ordenar(numbers []int) []int {
	sort.Slice(numbers, func(i, q int) bool {
		return numbers[i] < numbers[q]
	})
	return numbers
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return num / den, nil
}
