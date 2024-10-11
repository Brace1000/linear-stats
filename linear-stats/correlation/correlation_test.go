package  correlation 

import (
	"testing"
)

func TestPearson(t *testing.T) {
	
	correlation :=  CalculatePearsonCorrelation([]float64{1.0,6.0,8.0})

	expected := 0.9707253433941508 
	if correlation != expected {
		t.Errorf("correlation %v expected %f",correlation,expected)
	}
}
