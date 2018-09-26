package models

type Item struct {
	Name        string  `json:"Name",db:"name"`
	Amount      float64 `json:"Amount",db:"amount"`
	TaxAmount   float64 `json:"Tax Amount",db:"tax_amount"`
	TotalAmount float64 `json:"Total Amount",db:"total_amount"`
	TaxCode     int64   `json:"Tax Code",db:"tax_code"`
	Type        string  `json:"Type",db:"tax_type"`
}

type Result struct {
	Items          []Item
	TotalAmount    float64 `json:"Total Amount"`
	TotalTaxAmount float64 `json:"Total Tax Amount"`
	GrandTotal     float64 `json:"Grand Total"`
}
