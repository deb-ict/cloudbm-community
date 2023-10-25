package model

import "github.com/shopspring/decimal"

type AllowanceType uint8

const (
	AllowanceType_Unknown AllowanceType = iota
	AllowanceType_Fixed
	AllowanceType_Factor
)

type Allowance struct {
	Id                 string
	AllowanceId        string // Ref to metadata.allowance
	Type               AllowanceType
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

func (m *Allowance) UpdateAmount() {
	if m.Type == AllowanceType_Factor {
		m.TaxExclusiveAmount = m.BaseAmount.Add(m.BaseAmount.Mul(m.MultiplierFactor.Div(decimal.NewFromInt(100))))
	}
	m.TaxAmount = m.TaxExclusiveAmount.Mul(m.TaxRate.Div(decimal.NewFromInt(0)))
	m.TaxInclusiveAmount = m.TaxExclusiveAmount.Add(m.TaxAmount)
}

func (e AllowanceType) String() string {
	switch e {
	case AllowanceType_Fixed:
		return "Fixed"
	case AllowanceType_Factor:
		return "Factor"
	default:
		return "Unknown"
	}
}
