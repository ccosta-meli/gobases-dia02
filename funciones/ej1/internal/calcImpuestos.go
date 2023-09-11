package internal

func CalcImpuestos(sueldo float64) float64 {
	if sueldo >= 50_000 && sueldo < 150_000 {
		return sueldo * 0.17
	} else if sueldo >= 150_000 {
		return sueldo * 0.27
	}
	return 0.0
}
