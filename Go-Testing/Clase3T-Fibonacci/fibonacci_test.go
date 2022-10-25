package clase3tfibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonnaci(t *testing.T) {
	//Salida primeros 10 números
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34 ...
	fiboExpected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	fiboResult, err := Fibonacci(len(fiboExpected))

	assert.Nil(t, err)
	assert.Equal(t, fiboExpected, fiboResult)

	fiboResult, err = Fibonacci(-5)
	assert.NotNil(t, err)
	assert.Nil(t, fiboResult)

}
