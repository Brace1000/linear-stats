package regreation

// Calculating the LinearRegression
func CalculateLinearRegression(y_values []float64) (float64, float64) {
	n := float64(len(y_values))
	var sumX, sumY, sumXY, sumX2 float64
	for i, y := range y_values {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	/* calculating the slope m and intercept c */
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	c := (sumY - m*sumX) / n
	return m, c
}
