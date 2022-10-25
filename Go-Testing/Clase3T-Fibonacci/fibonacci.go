package clase3tfibonacci

import "errors"

func Fibonacci(n int) (result []int, err error) {
	if n < 0 {
		return result, errors.New("Negative numbers cannot be calculated")
	}
	f0, f1 := 0, 1
	for i := 0; i < n; i++ {
		result = append(result, f0)
		f0 = f1
		f1 = result[i] + f0
	}
	return result, nil
}
