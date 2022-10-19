package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1, num2 := 15, 20
	esperado := 5

	resultado := Restar(num1, num2)
	if resultado != esperado {
		t.Errorf("Función Restar() arrojó un resultado diferente al esperado")
	}
}

func TestRestarTestify(t *testing.T) {
	num1, num2 := 15, 20
	esperado := 5

	resultado := Restar(num1, num2)

	assert.Equal(t, esperado, resultado, "Deben ser iguales")
}

func TestOrdenar(t *testing.T) {
	numeros := []int{2, 5, 1, 4, 3}
	esperado := []int{1, 2, 3, 4, 5}

	resultado := Ordenar(numeros)

	assert.Equal(t, esperado, resultado, "Deben ser iguales")
}

func TestDividir(t *testing.T) {
	num := 100
	den := 0
	_, err := Dividir(num, den)
	assert.NotNil(t, err, "No llega al error")
}
