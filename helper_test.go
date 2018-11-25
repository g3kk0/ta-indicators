package taindicators

import (
	"fmt"
	"testing"
)

func TestAvg(t *testing.T) {
	res := Avg([]float64{1.1, 2.2, 3.3})

	s := fmt.Sprintf("%.1f", res)

	if s != "2.2" {
		t.Errorf("Avg was incorrect, got: %f, expected: %f.", res, 2.2)
	}
}

func TestSum(t *testing.T) {
	res := Sum([]float64{1.1, 2.2, 3.3})

	if res != 6.6 {
		t.Errorf("Sum was incorrect, got: %f, expected: %f.", res, 6.6)
	}
}
