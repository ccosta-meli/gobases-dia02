/*
02 - Calcular promedio.

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar **N** cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.

*/

package main

import (
	"dia-2/funciones/ej2/internal"
	"fmt"
)

func main() {
	fmt.Println("Las notas son: 7, 7, 7")
	fmt.Println("El promedio es:", internal.CalcPromedio(7, 7, 7))

	fmt.Println("Las notas son: 9, 8, 7")
	fmt.Println("El promedio es:", internal.CalcPromedio(9, 8, 7))

	fmt.Println("Las notas son: 9, 9, 10")
	fmt.Println("El promedio es:", internal.CalcPromedio(9, 9, 10))

	fmt.Println("Las notas son: 10, 10, 10")
	fmt.Println("El promedio es:", internal.CalcPromedio(10, 10, 10))
}
