/*
04 - Calcular estadísticas.

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
*/

package main

import (
	"dia-2/funciones/ej4/internal"
	"fmt"
)

func main() {
	var operation = "Promedio"
	var funcToCalc = internal.Operation(operation)
	var result = funcToCalc(7, 8, 9, 10)
	fmt.Println(operation, result)
}
