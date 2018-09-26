package src

import (
	"github.com/ramb0111/tax-calc/src/config"
)

type Payload struct {
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	TaxCode int16   `json:"taxcode"`
}

func (pl *Payload) foodTax() float64 {
	return 0.1 * pl.Amount
}

func (pl *Payload) tobaccoTax() float64 {
	return 10 + 0.02*pl.Amount
}

func (pl *Payload) entertainmentTax() float64 {
	if pl.Amount < 100 {
		return 0
	}
	return 0.01 * (pl.Amount - 100)
}

func (pl *Payload) GetTaxAmount() float64 {
	if pl.TaxCode == config.FoodTaxCode {
		return pl.foodTax()
	} else if pl.TaxCode == config.TobaccoTaxCode {
		return pl.tobaccoTax()
	} else if pl.TaxCode == config.EntertainmentTaxCode {
		return pl.entertainmentTax()
	}
	return -1
}
