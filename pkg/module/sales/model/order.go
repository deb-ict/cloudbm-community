package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	Number               string
	Date                 time.Time
	Recipient            *Party
	Delivery             *Party
	Items                []*OrderItem
	Allowances           []*Allowance
	Charges              []*Charge
	TaxAmounts           []*TaxAmount
	LineExtensionAmount  decimal.Decimal
	TaxExclusiveAmount   decimal.Decimal
	TaxInclusiveAmount   decimal.Decimal
	TaxTotalAmount       decimal.Decimal
	AllowanceTotalAmount decimal.Decimal
	ChargeTotalAmount    decimal.Decimal
	PrepaidAmount        decimal.Decimal
	PayableAmount        decimal.Decimal
}

func (m *Order) UpdateAmounts() {
	// Calculate the net amount
	m.LineExtensionAmount = decimal.Zero
	for _, item := range m.Items {
		item.UpdateAmounts()

		taxAmount := m.GetTax(item.TaxProfileId, item.TaxRate)
		taxAmount.BaseAmount.Add(item.TaxExclusiveAmount)

		m.LineExtensionAmount = m.LineExtensionAmount.Add(item.TaxExclusiveAmount)
	}

	// Calculate the allowance total amount
	m.AllowanceTotalAmount = decimal.Zero
	for _, allowance := range m.Allowances {
		allowance.UpdateAmount()

		taxAmount := m.GetTax(allowance.TaxProfileId, allowance.TaxRate)
		taxAmount.BaseAmount.Add(allowance.TaxExclusiveAmount)

		m.AllowanceTotalAmount.Add(allowance.TaxExclusiveAmount)
	}

	// Calculate the charges total amount
	m.ChargeTotalAmount = decimal.Zero
	for _, charge := range m.Charges {
		charge.UpdateAmounts()

		taxAmount := m.GetTax(charge.TaxProfileId, charge.TaxRate)
		taxAmount.BaseAmount.Sub(charge.TaxExclusiveAmount)

		m.ChargeTotalAmount.Add(charge.TaxExclusiveAmount)
	}

	// Calculate the tax amounts
	m.TaxExclusiveAmount = decimal.Zero
	m.TaxTotalAmount = decimal.Zero
	for _, tax := range m.TaxAmounts {
		tax.UpdateAmounts()

		m.TaxExclusiveAmount = m.TaxExclusiveAmount.Add(tax.BaseAmount)
		m.TaxTotalAmount = m.TaxTotalAmount.Add(tax.TaxAmount)
	}

	// Final total
	m.TaxInclusiveAmount = m.TaxExclusiveAmount.Add(m.TaxTotalAmount)
}

func (m *Order) GetTax(taxProfileId string, taxRate decimal.Decimal) *TaxAmount {
	for _, tax := range m.TaxAmounts {
		if tax.TaxProfileId == taxProfileId {
			return tax
		}
	}

	tax := &TaxAmount{
		TaxProfileId: taxProfileId,
		TaxRate:      taxRate,
		TaxAmount:    decimal.Zero,
		BaseAmount:   decimal.Zero,
	}
	m.TaxAmounts = append(m.TaxAmounts, tax)
	return tax
}
