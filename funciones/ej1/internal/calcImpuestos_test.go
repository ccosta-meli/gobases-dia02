package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcImpuestos(t *testing.T) {
	//Arrange -> Lo que necesita la funcion previo a que sea llamada para ejecutarse
	var sueldo float64 = 10_000
	var sueldo2 float64 = 70_000
	var sueldo3 float64 = 280_000

	//Act -> Llamado a la funcion
	impuesto1 := CalcImpuestos(sueldo)
	impuesto2 := CalcImpuestos(sueldo2)
	impuesto3 := CalcImpuestos(sueldo3)

	//Assert -> Comparacion entre lo esperado y lo que efectivamente sucedio.

	// 1. Esperamos que 10_000 no tenga impuestos, deberia retornar 0
	assert.Equal(t, 1.0, impuesto1, "Impuesto es 0")

	// Con testing sin complementos seria asi
	if impuesto1 != 0.0 {
		t.Errorf("La funcion CalcImp() arrojo un resultado distinto al esperado")
	}

	// 2. Esperamos que 70_000 tenga un impuesto del 17%, deberia retornar 11900.00
	assert.Equal(t, 11900.00, impuesto2, "Impuesto es 11900")

	// 3. Esperamos que 280_000 tenga un impuesto del 27%, deberia retornar 75600.00
	assert.Equal(t, 75600.00, impuesto3, "Impuesto es 75600")

	// Metodos mas utilizados
	// assert.Equal
	// assert.Nil
	// assert.Error
}
