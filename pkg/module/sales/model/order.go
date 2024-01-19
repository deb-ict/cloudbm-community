package model

import (
	"time"

	"github.com/shopspring/decimal"
)

/*
Order detail
-> show related invoices
-> show related shipments
*/

/*
OrderStatus
- Stage: PreOrder
	- Draft
	- PendingPayment
- Stage: Processing
	- Processing
	- OnHold
-Stage:Ended
	- Completed
	- Cancelled
	- Refunded
	- Failed
*/

type Order struct {
	Id                   string
	Number               string
	Date                 time.Time
	Recipient            *Party
	Delivery             *Party
	Items                []*OrderItem
	Allowances           []*DocumentAllowance
	Charges              []*DocumentCharge
	TaxAmounts           []*TaxAmount
	LineExtensionAmount  decimal.Decimal
	AllowanceTotalAmount decimal.Decimal
	ChargeTotalAmount    decimal.Decimal
	TaxExclusiveAmount   decimal.Decimal
	TaxInclusiveAmount   decimal.Decimal
	TaxTotalAmount       decimal.Decimal
}

type OrderFilter struct {
	Number  string
	MinDate time.Time
	MaxDate time.Time
}

func (m *Order) UpdateModel(other *Order) {
	m.Recipient = other.Recipient.Clone()
	m.Delivery = other.Delivery.Clone()
}

func (m *Order) UpdateAmounts() {
	// Reset the tax amounts
	m.TaxAmounts = make([]*TaxAmount, 0)

	// Calculate the net amount
	m.LineExtensionAmount = decimal.Zero
	for _, item := range m.Items {
		item.UpdateAmounts()
		taxAmount := m.GetTaxAmount(item.TaxRate)
		taxAmount.BaseAmount = taxAmount.BaseAmount.Add(item.LineTotal)
		m.LineExtensionAmount = m.LineExtensionAmount.Add(item.LineTotal)
	}

	// Calculate the allowance total amount
	m.AllowanceTotalAmount = decimal.Zero
	for _, allowance := range m.Allowances {
		allowance.UpdateAmounts()
		taxAmount := m.GetTaxAmount(allowance.TaxRate)
		taxAmount.BaseAmount = taxAmount.BaseAmount.Sub(allowance.Amount)
		m.AllowanceTotalAmount = m.AllowanceTotalAmount.Add(allowance.Amount)
	}

	// Calculate the charge total amount
	m.ChargeTotalAmount = decimal.Zero
	for _, charge := range m.Charges {
		charge.UpdateAmounts()
		taxAmount := m.GetTaxAmount(charge.TaxRate)
		taxAmount.BaseAmount = taxAmount.BaseAmount.Add(charge.Amount)
		m.ChargeTotalAmount = m.ChargeTotalAmount.Add(charge.Amount)
	}

	// Calculate the tax amounts
	m.TaxTotalAmount = decimal.Zero
	for _, taxAmount := range m.TaxAmounts {
		taxAmount.UpdateAmounts()
		m.TaxTotalAmount = m.TaxTotalAmount.Add(taxAmount.TaxAmount)
	}

	// Calculate total amount
	m.TaxExclusiveAmount = m.LineExtensionAmount.Add(m.AllowanceTotalAmount).Sub(m.ChargeTotalAmount)
	m.TaxInclusiveAmount = m.TaxExclusiveAmount.Add(m.TaxTotalAmount)
}

func (m *Order) GetTaxAmount(taxRate decimal.Decimal) *TaxAmount {
	for _, tax := range m.TaxAmounts {
		if tax.TaxRate.Equal(taxRate) {
			return tax
		}
	}
	tax := &TaxAmount{
		TaxRate:    taxRate,
		BaseAmount: decimal.Zero,
		TaxAmount:  decimal.Zero,
	}
	m.TaxAmounts = append(m.TaxAmounts, tax)
	return tax
}

func (m *Order) Clone() *Order {
	if m == nil {
		return nil
	}
	model := &Order{
		Id:                   m.Id,
		Number:               m.Number,
		Date:                 m.Date,
		Recipient:            m.Recipient.Clone(),
		Delivery:             m.Delivery.Clone(),
		Items:                make([]*OrderItem, 0),
		Allowances:           make([]*DocumentAllowance, 0),
		Charges:              make([]*DocumentCharge, 0),
		TaxAmounts:           make([]*TaxAmount, 0),
		LineExtensionAmount:  m.LineExtensionAmount,
		AllowanceTotalAmount: m.AllowanceTotalAmount,
		ChargeTotalAmount:    m.ChargeTotalAmount,
		TaxExclusiveAmount:   m.TaxExclusiveAmount,
		TaxInclusiveAmount:   m.TaxInclusiveAmount,
		TaxTotalAmount:       m.TaxTotalAmount,
	}
	for _, item := range m.Items {
		model.Items = append(model.Items, item.Clone())
	}
	for _, allowance := range m.Allowances {
		model.Allowances = append(model.Allowances, allowance.Clone())
	}
	for _, charge := range m.Charges {
		model.Charges = append(model.Charges, charge.Clone())
	}
	for _, taxAmount := range m.TaxAmounts {
		model.TaxAmounts = append(model.TaxAmounts, taxAmount.Clone())
	}
	return model
}
