package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcSalario(t *testing.T) {
	// Arrange
	minutosCatA := 4800
	minutosCatB := 4800
	minutosCatC := 4800

	// Act
	salarioA := CalcSalario(minutosCatA, "A")
	salarioB := CalcSalario(minutosCatB, "B")
	salarioC := CalcSalario(minutosCatC, "C")

	// Assert
	assert.Equal(t, 360_000.00, salarioA, "Salario categoria A deberia ser 360,000.00")
	assert.Equal(t, 144_000.00, salarioB, "Salario categoria B deberia ser 144,000.00")
	assert.Equal(t, 80_000.00, salarioC, "Salario categoria C deberia ser 80,000.00")

}
