package model

import "github.com/shopspring/decimal"

type TaxAmount struct {
	TaxProfileId string
	TaxRate      decimal.Decimal
	TaxAmount    decimal.Decimal
	BaseAmount   decimal.Decimal
}

func (m *TaxAmount) UpdateAmounts() {
	m.TaxAmount = m.BaseAmount.Mul(m.TaxRate.Div(decimal.NewFromInt(100)))
}
