package model

import (
	"github.com/shopspring/decimal"
)

type InvoiceLine struct {
	Id                 string
	OrderItemId        string // ref to order.line.id
	ArticleType        ArticleType
	ArticleId          string
	ArticleReference   string
	Description        string
	TaxProfileId       string // ref to metadata.taxProfile
	TaxRate            decimal.Decimal
	Quantity           decimal.Decimal
	UnitPrice          decimal.Decimal
	TaxExclusiveAmount decimal.Decimal
	TaxInclusiveAmount decimal.Decimal
	TaxAmount          decimal.Decimal
}

func (m *InvoiceLine) UpdateAmounts() {
	m.TaxExclusiveAmount = m.UnitPrice.Mul(m.Quantity)
	m.TaxAmount = m.TaxExclusiveAmount.Mul(m.TaxRate.Div(decimal.NewFromInt(100)))
	m.TaxInclusiveAmount = m.TaxExclusiveAmount.Add(m.TaxAmount)
}
