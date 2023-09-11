package internal

func calcularMaximo(notas ...int) int {
	if len(notas) == 0 {
		return 0
	}
	max := notas[0]

	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return max
}

func calcularMinimo(notas ...int) int {
	if len(notas) == 0 {
		return 0
	}
	min := notas[0]

	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return min
}

func calcularPromedio(notas ...int) int {
	suma := 0
	for _, nota := range notas {
		suma += nota
	}
	return suma / len(notas)
}

func Operation(oper string) func(notas ...int) int {
	switch oper {
	case "Maximo":
		return calcularMaximo
	case "Minimo":
		return calcularMinimo
	case "Promedio":
		return calcularPromedio
	default:
		return nil
	}
}
