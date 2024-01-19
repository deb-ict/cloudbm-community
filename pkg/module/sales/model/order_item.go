package model

import "github.com/shopspring/decimal"

type OrderItem struct {
	Id          string
	ArticleType ArticleType
	ArticleId   string
	Description string
	Quantity    decimal.Decimal
	UnitPrice   decimal.Decimal
	LineTotal   decimal.Decimal
	Allowances  []*ItemAllowance
	Charges     []*ItemCharge
	TaxRate     decimal.Decimal
}

type OrderItemFilter struct {
}

func (m *OrderItem) UpdateAmounts() {
	m.LineTotal = m.UnitPrice.Mul(m.Quantity)
	for _, allowance := range m.Allowances {
		allowance.UpdateAmounts()
		m.LineTotal = m.LineTotal.Sub(allowance.Amount)
	}
	for _, charge := range m.Charges {
		charge.UpdateAmounts()
		m.LineTotal = m.LineTotal.Add(charge.Amount)
	}
}

func (m *OrderItem) Clone() *OrderItem {
	model := &OrderItem{
		Id:          m.Id,
		ArticleType: m.ArticleType,
		ArticleId:   m.ArticleId,
		Description: m.Description,
		Quantity:    m.Quantity,
		UnitPrice:   m.UnitPrice,
		LineTotal:   m.LineTotal,
		Allowances:  make([]*ItemAllowance, 0),
		Charges:     make([]*ItemCharge, 0),
		TaxRate:     m.TaxRate,
	}
	for _, allowance := range m.Allowances {
		model.Allowances = append(model.Allowances, allowance.Clone())
	}
	for _, charge := range m.Charges {
		model.Charges = append(model.Charges, charge.Clone())
	}
	return model
}
