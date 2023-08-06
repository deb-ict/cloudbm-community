package model

import "time"

type PurchaseInvoice struct {
	Id            string
	SupplierId    string
	InvoiceNumber string
	InvoiceDate   time.Time
	DueDate       time.Time
	Lines         []*PurchaseInvoiceLine
}
