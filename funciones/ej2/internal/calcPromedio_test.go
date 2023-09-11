package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPromedio(t *testing.T) {
	// Arrange

	// Act
	promedio := CalcPromedio(7, 7, 8, 9)

	// Assert
	assert.Equal(t, 7.75, promedio, "Test deberia dar 7.75")
}
