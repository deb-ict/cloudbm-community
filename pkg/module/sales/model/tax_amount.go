package model

import "github.com/shopspring/decimal"

type TaxAmount struct {
	BaseAmount decimal.Decimal
	TaxRate    decimal.Decimal
	TaxAmount  decimal.Decimal
}

func (m *TaxAmount) UpdateAmounts() {
	m.TaxAmount = m.BaseAmount.Mul(m.TaxRate.Div(decimal.NewFromInt(100)))
}

func (m *TaxAmount) Clone() *TaxAmount {
	return &TaxAmount{
		BaseAmount: m.BaseAmount,
		TaxRate:    m.TaxRate,
		TaxAmount:  m.TaxAmount,
	}
}
