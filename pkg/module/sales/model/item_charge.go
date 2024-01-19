package model

import "github.com/shopspring/decimal"

type ItemCharge struct {
	Id               string
	Type             ChargeType
	ReasonCode       string
	Reason           string
	Amount           decimal.Decimal
	BaseAmount       decimal.Decimal
	MultiplierFactor decimal.Decimal
}

func (m *ItemCharge) UpdateAmounts() {
	if m.Type == ChargeType_Factor {
		m.Amount = m.BaseAmount.Add(m.BaseAmount.Mul(m.MultiplierFactor.Div(decimal.NewFromInt(100))))
	}
}

func (m *ItemCharge) Clone() *ItemCharge {
	return &ItemCharge{
		Id:               m.Id,
		Type:             m.Type,
		ReasonCode:       m.ReasonCode,
		Reason:           m.Reason,
		Amount:           m.Amount,
		BaseAmount:       m.BaseAmount,
		MultiplierFactor: m.MultiplierFactor,
	}
}
