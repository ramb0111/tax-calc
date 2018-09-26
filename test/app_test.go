package src

import (
	"testing"

	"github.com/ramb0111/tax-calc/src"
)

func TestTax(t *testing.T) {
	pls := []src.Payload{
		{"tax1", 1000, 1},
		{"tax2", 2000, 2},
		{"tax3", 3000, 3},
	}

	amt := [3]float64{100, 50, 29}

	for i, pl := range pls {
		if amt[i] != pl.GetTaxAmount() {
			t.Errorf("Tax for payload name %s is %f, while it must be %f ", pl.Name, pl.GetTaxAmount(), amt[i])
		}
	}
}
