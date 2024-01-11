package model

import "github.com/shopspring/decimal"

type DocumentCharge struct {
	Id               string
	Type             ChargeType
	ReasonCode       string
	Reason           string
	Amount           decimal.Decimal
	BaseAmount       decimal.Decimal
	MultiplierFactor decimal.Decimal
	TaxRate          decimal.Decimal
}

func (m *DocumentCharge) UpdateAmounts() {
	if m.Type == ChargeType_Factor {
		m.Amount = m.BaseAmount.Add(m.BaseAmount.Mul(m.MultiplierFactor.Div(decimal.NewFromInt(100))))
	}
}

func (m *DocumentCharge) Clone() *DocumentCharge {
	return &DocumentCharge{
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
