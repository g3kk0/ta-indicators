package taindicators

import (
	"testing"
)

func TestSma(t *testing.T) {
	res := Sma([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 3)

	//	if res != 6.6 {
	//		t.Errorf("Sum was incorrect, got: %f, expected: %f.", res, 6.6)
	//	}
}
