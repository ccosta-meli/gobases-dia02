package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularMaximo(t *testing.T) {
	maximo := calcularMaximo(1, 2, 6, 8, 4, 8, 9)
	assert.Equal(t, 9, maximo, "Maximo deberia ser 9")
}

func TestCalcularMinimo(t *testing.T) {
	minimo := calcularMinimo(1, 2, 6, 8, 4, 8, 9)
	assert.Equal(t, 1, minimo, "Minimo deberia ser 1")
}

func TestCalcularPromedio(t *testing.T) {
	promedio := calcularPromedio(7, 7, 8, 9, 10)
	assert.Equal(t, 8, promedio, "Promedio deberia ser 8")
}
