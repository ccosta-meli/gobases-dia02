package internal

func CalcSalario(minutos int, cat string) float64 {
	horas := minutos / 60
	switch cat {
	case "A":
		return (float64(horas*3000) * 1.50)
	case "B":
		return (float64(horas*1500) * 1.20)
	default:
		return (float64(horas*1000) * 1)
	}
}
