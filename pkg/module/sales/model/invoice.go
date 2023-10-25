package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	Id                      string
	Number                  string
	IssueDate               time.Time
	DueDate                 time.Time
	OrderId                 string
	OrderNumber             string
	PaymentTermsId          string
	PaymentTermsDescription string
	Recipient               *Party
	Delivery                *Party
	Lines                   []*InvoiceLine
	Allowances              []*Allowance
	Charges                 []*Charge
	TaxAmounts              []*TaxAmount
	LineExtensionAmount     decimal.Decimal
	TaxExclusiveAmount      decimal.Decimal
	TaxInclusiveAmount      decimal.Decimal
	TaxTotalAmount          decimal.Decimal
	AllowanceTotalAmount    decimal.Decimal
	ChargeTotalAmount       decimal.Decimal
	PrepaidAmount           decimal.Decimal
	PayableAmount           decimal.Decimal
}

func (m *Invoice) UpdateAmounts() {
	// Calculate the net amount
	m.LineExtensionAmount = decimal.Zero
	for _, line := range m.Lines {
		line.UpdateAmounts()

		taxAmount := m.GetTax(line.TaxProfileId, line.TaxRate)
		taxAmount.BaseAmount.Add(line.TaxExclusiveAmount)

		m.LineExtensionAmount = m.LineExtensionAmount.Add(line.TaxExclusiveAmount)
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

func (m *Invoice) GetTax(taxProfileId string, taxRate decimal.Decimal) *TaxAmount {
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
