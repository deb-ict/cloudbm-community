package model

import "github.com/shopspring/decimal"

type ItemAllowance struct {
	Id               string
	Type             AllowanceType
	ReasonCode       string
	Reason           string
	Amount           decimal.Decimal
	BaseAmount       decimal.Decimal
	MultiplierFactor decimal.Decimal
}

func (m *ItemAllowance) UpdateAmounts() {
	if m.Type == AllowanceType_Factor {
		m.Amount = m.BaseAmount.Add(m.BaseAmount.Mul(m.MultiplierFactor.Div(decimal.NewFromInt(100))))
	}
}

func (m *ItemAllowance) Clone() *ItemAllowance {
	return &ItemAllowance{
		Id:               m.Id,
		Type:             m.Type,
		ReasonCode:       m.ReasonCode,
		Reason:           m.Reason,
		Amount:           m.Amount,
		BaseAmount:       m.BaseAmount,
		MultiplierFactor: m.MultiplierFactor,
	}
}
