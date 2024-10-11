package correlation

import (
	"math"
)

// Calculate the pearson correlation from the data
func CalculatePearsonCorrelation(y_values []float64) float64 {
	n := float64(len(y_values))
	var sumX, sumX2, sumXY, sumY, sumY2 float64

	for i, y := range y_values {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}
	r := (n*sumXY - sumX*sumY) / math.Sqrt((n*sumX2-sumX*sumX)*(n*sumY2-sumY*sumY))
	return r
}
