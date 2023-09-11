package internal

func CalcPromedio(notas ...int) float64 {
	cantNotas := float64(len(notas))
	var suma float64 = 0
	for _, nota := range notas {
		suma += float64(nota)
	}
	return suma / cantNotas
}
