package agristat

func CalculateAverageYield(yields []float64) float64 {
	sum := 0.0
	for _, y := range yields {
		sum += y
	}
	if len(yields) == 0 {
		return 0
	}
	return sum / float64(len(yields))
}
