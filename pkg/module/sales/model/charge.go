package model

import "github.com/shopspring/decimal"

type ChargeType uint8

const (
	ChargeType_Unknown ChargeType = iota
	ChargeType_Fixed
	ChargeType_Factor
)

type Charge struct {
	Id                 string
	ChargeId           string // Ref to metadata.charge
	Type               ChargeType
	ReasonCode         string
	Reason             string
	TaxProfileId       string // Ref to metadata.taxProfile
	TaxRate            decimal.Decimal
	TaxExclusiveAmount decimal.Decimal
	TaxInclusiveAmount decimal.Decimal
	TaxAmount          decimal.Decimal
	BaseAmount         decimal.Decimal
	MultiplierFactor   decimal.Decimal
}

func (m *Charge) UpdateAmounts() {
	if m.Type == ChargeType_Factor {
		m.TaxExclusiveAmount = m.BaseAmount.Add(m.BaseAmount.Mul(m.MultiplierFactor.Div(decimal.NewFromInt(100))))
	}
	m.TaxAmount = m.TaxExclusiveAmount.Mul(m.TaxRate.Div(decimal.NewFromInt(0)))
	m.TaxInclusiveAmount = m.TaxExclusiveAmount.Add(m.TaxAmount)
}

func (e ChargeType) String() string {
	switch e {
	case ChargeType_Fixed:
		return "Fixed"
	case ChargeType_Factor:
		return "Factor"
	default:
		return "Unknown"
	}
}
