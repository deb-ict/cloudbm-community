package model

import "github.com/shopspring/decimal"

type DocumentAllowance struct {
	Id               string
	Type             AllowanceType
	ReasonCode       string
	Reason           string
	Amount           decimal.Decimal
	BaseAmount       decimal.Decimal
	MultiplierFactor decimal.Decimal
	TaxRate          decimal.Decimal
}

func (m *DocumentAllowance) UpdateAmounts() {
	if m.Type == AllowanceType_Factor {
		m.Amount = m.BaseAmount.Add(m.BaseAmount.Mul(m.MultiplierFactor.Div(decimal.NewFromInt(100))))
	}
}

func (m *DocumentAllowance) Clone() *DocumentAllowance {
	return &DocumentAllowance{
		Id:               m.Id,
		Type:             m.Type,
		ReasonCode:       m.ReasonCode,
		Reason:           m.Reason,
		Amount:           m.Amount,
		BaseAmount:       m.BaseAmount,
		MultiplierFactor: m.MultiplierFactor,
		TaxRate:          m.TaxRate,
	}
}
