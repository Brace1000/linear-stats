package regreation

import (
	"testing"
)

func TestCalculateLinearRegression(t *testing.T) {
	// Sample y_values for testing
	y_values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	
	// Expected results for slope (m) and intercept (c)
	expectedM := 1.0 // Slope
	expectedC := 1.0 // Intercept

	// Call the function to calculate linear regression
	m, c := CalculateLinearRegression(y_values)

	// Check if the calculated slope is as expected
	if m != expectedM {
		t.Errorf("expected slope %f, got %f", expectedM, m)
	}

	// Check if the calculated intercept is as expected
	if c != expectedC {
		t.Errorf("expected intercept %f, got %f", expectedC, c)
	}
}
